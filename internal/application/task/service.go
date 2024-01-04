// Package task provides
package task

type Service struct {
	postgresRepo Repository
}

// ServiceParam .
type ServiceParam struct {
	PostgresRepo Repository
}

// NewService .
func NewService(param ServiceParam) *Service {
	return &Service{
		postgresRepo: param.PostgresRepo,
	}
}
