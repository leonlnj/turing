// Code generated by mockery v2.6.0. DO NOT EDIT.

package mocks

import (
	imagebuilder "github.com/gojek/turing/api/turing/imagebuilder"
	mock "github.com/stretchr/testify/mock"

	models "github.com/gojek/turing/api/turing/models"
)

// ImageBuilder is an autogenerated mock type for the ImageBuilder type
type ImageBuilder struct {
	mock.Mock
}

// BuildImage provides a mock function with given fields: request
func (_m *ImageBuilder) BuildImage(request imagebuilder.BuildImageRequest) (string, error) {
	ret := _m.Called(request)

	var r0 string
	if rf, ok := ret.Get(0).(func(imagebuilder.BuildImageRequest) string); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(imagebuilder.BuildImageRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteImageBuildingJob provides a mock function with given fields: projectName, modelName, versionID
func (_m *ImageBuilder) DeleteImageBuildingJob(projectName string, modelName string, versionID models.ID) error {
	ret := _m.Called(projectName, modelName, versionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, models.ID) error); ok {
		r0 = rf(projectName, modelName, versionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetImageBuildingJobStatus provides a mock function with given fields: projectName, modelName, versionID
func (_m *ImageBuilder) GetImageBuildingJobStatus(projectName string, modelName string, versionID models.ID) (imagebuilder.JobStatus, error) {
	ret := _m.Called(projectName, modelName, versionID)

	var r0 imagebuilder.JobStatus
	if rf, ok := ret.Get(0).(func(string, string, models.ID) imagebuilder.JobStatus); ok {
		r0 = rf(projectName, modelName, versionID)
	} else {
		r0 = ret.Get(0).(imagebuilder.JobStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, models.ID) error); ok {
		r1 = rf(projectName, modelName, versionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
