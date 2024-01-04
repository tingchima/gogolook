// Package application provides
package application

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/tingchima/gogolook/internal/application/task"
	"github.com/tingchima/gogolook/internal/repository/postgres"
)

// Application .
type Application struct {
	TaskService *task.Service
}

// ApplicationParam .
type ApplicationParam struct {
	PostgresConn *sqlx.DB
}

// MustNewApplication .
func MustNewApplication(param ApplicationParam) *Application {

	app, err := NewApplication(param)
	if err != nil {
		log.Panicf("new application fail, err: %s", err.Error())
		return nil
	}

	return app
}

// NewApplication .
func NewApplication(param ApplicationParam) (*Application, error) {

	postgresRepo := postgres.NewRepository(param.PostgresConn)

	taskService := task.NewService(task.ServiceParam{
		PostgresRepo: postgresRepo,
	})

	return &Application{TaskService: taskService}, nil
}
