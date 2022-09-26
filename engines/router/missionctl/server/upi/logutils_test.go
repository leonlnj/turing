package upi

import (
	"context"
	"fmt"
	"testing"

	"github.com/caraml-dev/turing/engines/router/missionctl/errors"
	upiv1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	fiberProtocol "github.com/gojek/fiber/protocol"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func Test_logTuringRouterRequestSummary(t *testing.T) {
	resp := &upiv1.PredictValuesResponse{}
	respByte, err := proto.Marshal(resp)
	require.NoError(t, err)
	key := "test"

	tests := []struct {
		name     string
		err      *errors.TuringError
		expected grpcRouterResponse
	}{
		{
			name: "ok",
			expected: grpcRouterResponse{
				key:  key,
				body: respByte,
			},
		},
		{
			name: "error",
			err:  errors.NewTuringError(fmt.Errorf("test error"), fiberProtocol.GRPC),
			expected: grpcRouterResponse{
				key: key,
				err: "test error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctx := metadata.AppendToOutgoingContext(context.Background(), "test", "key")
			// Make response channel
			respCh := make(chan grpcRouterResponse, 1)
			copyResponseToLogChannel(ctx, respCh, key, respByte, tt.err)

			close(respCh)
			data := <-respCh
			require.Equal(t, tt.expected, data)
		})
	}
}
