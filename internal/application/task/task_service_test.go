// Package task provides
package task

import (
	"context"
	"errors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tingchima/gogolook/internal/domain"
	"github.com/tingchima/gogolook/internal/domain/common"
)

// TestTaskService_ListTasks .
func TestTaskService_ListTasks(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type Args struct {
		Tasks []domain.Task `faker:"-"`
	}

	var args Args

	err := faker.FakeData(&args)
	require.NoError(t, err)

	args.Tasks = make([]domain.Task, 20)
	for i := range args.Tasks {
		err := faker.FakeData(&args.Tasks[i])
		require.NoError(t, err)
	}

	tests := []struct {
		name            string
		wantErr         bool
		expectedErrCode common.ErrCode
		setupService    func(t *testing.T) *Service
	}{
		{
			name: "success",
			setupService: func(t *testing.T) *Service {
				mock := buildMockService(ctrl)

				mock.postgresRepo.EXPECT().ListTasks(gomock.Any(), gomock.Any()).Return(args.Tasks, nil)

				return buildService(mock)
			},
			wantErr: false,
		},
		{
			name: "internal server error",
			setupService: func(t *testing.T) *Service {
				mock := buildMockService(ctrl)

				err := common.NewError(common.ErrCodeInternalProcess, errors.New("mock db server error"))

				mock.postgresRepo.EXPECT().ListTasks(gomock.Any(), gomock.Any()).Return(nil, err)

				return buildService(mock)
			},
			wantErr:         true,
			expectedErrCode: common.ErrCodeInternalProcess,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.setupService(t)

			param := domain.TaskParam{}

			_, err := s.ListTasks(context.Background(), param)
			if tt.wantErr {
				require.Error(t, err)

				var domainErr *common.Error
				assert.True(t, common.AsErr(err, &domainErr))
				assert.True(t, common.IsErrCode(err, tt.expectedErrCode))

			} else {
				require.NoError(t, err)
			}
		})
	}
}

// TestTaskService_CreateTask .
func TestTaskService_CreateTask(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type Args struct {
		Task domain.Task
	}

	var args Args

	err := faker.FakeData(&args)
	require.NoError(t, err)

	tests := []struct {
		name            string
		wantErr         bool
		expectedErrCode common.ErrCode
		setupService    func(t *testing.T) *Service
	}{
		{
			name: "success",
			setupService: func(t *testing.T) *Service {
				mock := buildMockService(ctrl)

				mock.postgresRepo.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Return(&args.Task, nil)

				return buildService(mock)
			},
			wantErr: false,
		},
		{
			name: "internal server error",
			setupService: func(t *testing.T) *Service {
				mock := buildMockService(ctrl)

				err := common.NewError(common.ErrCodeInternalProcess, errors.New("mock db server error"))

				mock.postgresRepo.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Return(nil, err)

				return buildService(mock)
			},
			wantErr:         true,
			expectedErrCode: common.ErrCodeInternalProcess,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.setupService(t)

			got, err := s.CreateTask(context.Background(), args.Task)
			if tt.wantErr {
				require.Error(t, err)

				var domainErr *common.Error
				assert.True(t, common.AsErr(err, &domainErr))
				assert.True(t, common.IsErrCode(err, tt.expectedErrCode))

			} else {
				require.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}

// TestTaskService_UpdateTask .
func TestTaskService_UpdateTask(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type Args struct {
		Task domain.Task
	}

	var args Args

	err := faker.FakeData(&args)
	require.NoError(t, err)

	tests := []struct {
		name            string
		wantErr         bool
		expectedErrCode common.ErrCode
		setupService    func(t *testing.T) *Service
	}{
		{
			name: "success",
			setupService: func(t *testing.T) *Service {
				mock := buildMockService(ctrl)

				mock.postgresRepo.EXPECT().UpdateTask(gomock.Any(), gomock.Any()).Return(&args.Task, nil)

				return buildService(mock)
			},
			wantErr: false,
		},
		{
			name: "task not found error",
			setupService: func(t *testing.T) *Service {
				mock := buildMockService(ctrl)

				err := common.NewError(common.ErrCodeResourceNotFound, errors.New("mock task not found error"))

				mock.postgresRepo.EXPECT().UpdateTask(gomock.Any(), gomock.Any()).Return(nil, err)

				return buildService(mock)
			},
			wantErr:         true,
			expectedErrCode: common.ErrCodeResourceNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.setupService(t)

			got, err := s.UpdateTask(context.Background(), args.Task)
			if tt.wantErr {
				require.Error(t, err)

				var domainErr *common.Error
				assert.True(t, common.AsErr(err, &domainErr))
				assert.True(t, common.IsErrCode(err, tt.expectedErrCode))

			} else {
				require.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}

// TestTaskService_DeleteTask .
func TestTaskService_DeleteTask(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type Args struct {
		TaskID int64
	}

	var args Args

	err := faker.FakeData(&args)
	require.NoError(t, err)

	tests := []struct {
		name            string
		wantErr         bool
		expectedErrCode common.ErrCode
		setupService    func(t *testing.T) *Service
	}{
		{
			name: "success",
			setupService: func(t *testing.T) *Service {
				mock := buildMockService(ctrl)

				mock.postgresRepo.EXPECT().DeleteTaskByID(gomock.Any(), gomock.Any()).Return(nil)

				return buildService(mock)
			},
			wantErr: false,
		},
		{
			name: "task not found error",
			setupService: func(t *testing.T) *Service {
				mock := buildMockService(ctrl)

				err := common.NewError(common.ErrCodeResourceNotFound, errors.New("mock task not found error"))

				mock.postgresRepo.EXPECT().DeleteTaskByID(gomock.Any(), gomock.Any()).Return(err)

				return buildService(mock)
			},
			wantErr:         true,
			expectedErrCode: common.ErrCodeResourceNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.setupService(t)

			err := s.DeleteTaskByID(context.Background(), args.TaskID)
			if tt.wantErr {
				require.Error(t, err)

				var domainErr *common.Error
				assert.True(t, common.AsErr(err, &domainErr))
				assert.True(t, common.IsErrCode(err, tt.expectedErrCode))

			} else {
				require.NoError(t, err)
			}
		})
	}
}
