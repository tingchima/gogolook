// Package main provides
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tingchima/gogolook/configs"
	"github.com/tingchima/gogolook/infra"
	"github.com/tingchima/gogolook/internal/application"
	handler_http "github.com/tingchima/gogolook/internal/handler/http"
)

const (
	// config file
	DefaultDBUsername = "postgres"
	DefaultDBPassword = "postgres"
	DefaultDBHost     = "localhost"
	DefaultDBPort     = "5432"
	DefaultDBName     = "gogolook"

	DefaultServerPort = "8080"
)

// @title gogolook api server
// @version 1.0
// @description api server
func main() {

	rootCtx, rootCtxCancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// run server
	RunServer(rootCtx, wg)

	// graceful shutdown process
	{
		shutdownCh := make(chan os.Signal)
		signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

		<-shutdownCh

		// execute root context cancel to shutdown
		rootCtxCancel()
	}

	// wait all process be done
	{
		wg.Wait()
	}
}

// RunServer .
func RunServer(rootCtx context.Context, wg *sync.WaitGroup) {

	// new relative infra
	postgresConn := infra.MustNewPostgresConn(&configs.Database{
		Username: DefaultDBUsername,
		Password: DefaultDBPassword,
		Host:     DefaultDBHost,
		DBName:   DefaultDBName,
		Port:     DefaultDBPort,
	})

	// new app
	app := application.MustNewApplication(application.ApplicationParam{
		PostgresConn: postgresConn,
	})

	// new handler
	handler := gin.New()

	// use middleware
	handler.Use(gin.Recovery())

	// register http handlers
	handler_http.RegisterHandlers(handler, app)

	// new server
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", DefaultServerPort),
		Handler: handler,
	}

	// run server
	{
		// run http server in background
		go func() {
			defer wg.Done()

			err := server.ListenAndServe()
			if err != nil {
				log.Println(err.Error())
			}
		}()

		// init shutdown http server process in background
		go func() {
			defer wg.Done()

			<-rootCtx.Done()

			shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			err := server.Shutdown(shutdownCtx)
			if err != nil {
				log.Println(err.Error())
			}
		}()
	}
}
