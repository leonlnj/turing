package upi

import (
	"bytes"
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caraml-dev/turing/engines/router/missionctl"
	"github.com/caraml-dev/turing/engines/router/missionctl/errors"
	"github.com/caraml-dev/turing/engines/router/missionctl/instrumentation/metrics"
	"github.com/caraml-dev/turing/engines/router/missionctl/instrumentation/tracing"
	"github.com/caraml-dev/turing/engines/router/missionctl/log"
	"github.com/caraml-dev/turing/engines/router/missionctl/log/resultlog"
	"github.com/caraml-dev/turing/engines/router/missionctl/server/constant"
	"github.com/caraml-dev/turing/engines/router/missionctl/turingctx"
	"github.com/gojek/fiber"
	fibergrpc "github.com/gojek/fiber/grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

const tracingComponentID = "grpc_handler"

type Server struct {
	missionControl missionctl.MissionControlUPI
}

func NewUPIServer(mc missionctl.MissionControlUPI) *Server {
	return &Server{
		missionControl: mc,
	}
}

// UniversalPredictionServiceServer is the internal server that communicates in bytes
// It requires the codec to be passing []byte instead of a proto object. It is registered
// the same way as the generated proto which has the following interface instead.
// PredictValues(context.Context, *PredictValuesRequest) (*PredictValuesResponse, error)
type UniversalPredictionServiceServer interface {
	PredictValues(context.Context, []byte) ([]byte, error)
}

//nolint:all
// UniversalPredictionServiceServiceDesc implement grpc internal methodHandler interface, it creates an io.writer
// and expect the codec to write bytes into the input instead of the typical proto request.
// Nolint is used as context cannot be changed to first argument in order to fit the interface.
func UniversalPredictionServiceServiceDesc(
	srv interface{},
	ctx context.Context,
	dec func(interface{}) error,
	interceptor grpc.UnaryServerInterceptor,
) (interface{}, error) {
	in := new(bytes.Buffer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversalPredictionServiceServer).PredictValues(ctx, in.Bytes())
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/caraml.upi.v1.UniversalPredictionService/PredictValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversalPredictionServiceServer).PredictValues(ctx, req.([]byte))
	}
	return interceptor(ctx, in.Bytes(), info, handler)
}

// ServiceDesc is the grpc.ServiceDesc for UniversalPredictionService service.
// It is and must register the same service as per the generated proto, in order
// for client to recognize it. However underlying it this is using a different handler
// which uses and expects bytes input instead of the proto message. Any form of validation
// of the input/output proto while calling this service is lost due to the byte message not
// being decoded or encoded by Turing at any point.
var ServiceDesc = grpc.ServiceDesc{
	ServiceName: "caraml.upi.v1.UniversalPredictionService",
	HandlerType: (*UniversalPredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PredictValues",
			Handler:    UniversalPredictionServiceServiceDesc,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "caraml/upi/v1/upi.proto",
}

func (us *Server) Run(listener net.Listener) {
	// ForceServerCodec is experimental API and fiber codec is used to
	// work with the server internal implementation to receive bytes without (un)marshalling to
	// minimize overhead while losing validation of the proto message.
	s := grpc.NewServer(grpc.ForceServerCodec(&fibergrpc.FiberCodec{}))
	s.RegisterService(&ServiceDesc, us)
	reflection.Register(s)

	errChan := make(chan error, 1)
	stopChan := make(chan os.Signal, 1)

	// bind OS events to the signal channel
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := s.Serve(listener); err != nil {
			errChan <- err
		}
	}()

	defer func() {
		s.GracefulStop()
	}()

	// block until either OS signal, or server fatal error
	select {
	case err := <-errChan:
		log.Glob().Errorf("Failed to start Turing Mission Control API: %s", err)
	case <-stopChan:
		log.Glob().Info("Signal to stop server")
	}

}

func (us *Server) PredictValues(ctx context.Context, req []byte) (
	[]byte, error) {
	var predictionErr *errors.TuringError // Measure execution time
	defer metrics.Glob().MeasureDurationMs(
		metrics.TuringComponentRequestDurationMs,
		map[string]func() string{
			"status": func() string {
				return metrics.GetStatusString(predictionErr == nil)
			},
			"component": func() string {
				return tracingComponentID
			},
		},
	)()

	// Create context from the request context
	ctx = turingctx.NewTuringContext(ctx)
	// Create context logger
	ctxLogger := log.WithContext(ctx)

	// if request comes with metadata, attach it to metadata to be sent with fiber
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(map[string]string{})
	}

	// Get the unique turing request id from the context
	turingReqID, err := turingctx.GetRequestID(ctx)
	if err != nil {
		ctxLogger.Errorf("Could not retrieve Turing Request ID from context: %v",
			err.Error())
	}
	ctxLogger.Debugf("Received request for %v", turingReqID)
	md.Append(constant.TuringReqIDHeaderKey, turingReqID)

	if tracing.Glob().IsEnabled() {
		var sp opentracing.Span
		sp, _ = tracing.Glob().StartSpanFromContext(ctx, tracingComponentID)
		if sp != nil {
			defer sp.Finish()
		}
	}

	fiberRequest := &fibergrpc.Request{
		Message:  req,
		Metadata: md,
	}
	resp, predictionErr := us.getPrediction(ctx, fiberRequest)
	if predictionErr != nil {
		logTuringRouterRequestError(ctx, predictionErr)
		return nil, predictionErr
	}
	return resp, nil
}

func (us *Server) getPrediction(
	ctx context.Context,
	fiberRequest fiber.Request) (
	[]byte, *errors.TuringError) {

	// Create response channel to store the response from each step. 1 for route now,
	// should be 4 when experiment engine, enricher and ensembler are added
	respCh := make(chan grpcRouterResponse, 1)

	// Defer logging request summary
	defer func() {
		go func() {
			logTuringRouterRequestSummary(
				ctx,
				time.Now(),
				fiberRequest.Header(),
				fiberRequest.Payload(),
				respCh)
			close(respCh)
		}()
	}()

	// Calling Routes via fiber
	resp, err := us.missionControl.Route(ctx, fiberRequest)
	if err != nil {
		return nil, err
	}
	copyResponseToLogChannel(ctx, respCh, resultlog.ResultLogKeys.Router, resp, err)

	return resp, nil
}
