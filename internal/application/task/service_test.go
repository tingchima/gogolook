// Package task provides
package task

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tingchima/gogolook/internal/application/task/mocks"
)

func TestMain(m *testing.M) {
	_ = m.Run()
}

// mockService .
type mockService struct {
	postgresRepo *mocks.MockRepository
}

// buildMockService .
func buildMockService(ctrl *gomock.Controller) mockService {

	return mockService{
		postgresRepo: mocks.NewMockRepository(ctrl),
	}
}

// buildService .
func buildService(param mockService) *Service {

	return NewService(ServiceParam{
		PostgresRepo: param.postgresRepo,
	})
}
