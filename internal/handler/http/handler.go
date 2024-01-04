// Package http provides
package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tingchima/gogolook/internal/application"
)

// RegisterHandlers .
func RegisterHandlers(handler *gin.Engine, app *application.Application) {

	// task handlers
	{
		handler.GET("/tasks", ListTasks(app))

		handler.POST("/task", CreateTask(app))

		handler.PUT("/task/:id", UpdateTask(app))

		handler.DELETE("/task/:id", DeleteTask(app))
	}
}
