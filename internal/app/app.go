// Package app configures and runs application.
package app

import (
	"fmt"

	"github.com/hyuti/pocketbase-template/cmd"
	"github.com/hyuti/pocketbase-template/config"
	"github.com/hyuti/pocketbase-template/internal/controller"
	"github.com/hyuti/pocketbase-template/internal/schedule"
	"github.com/hyuti/pocketbase-template/internal/task"
	"github.com/hyuti/pocketbase-template/internal/validation"
	"github.com/hyuti/pocketbase-template/internal/webapi"
	_ "github.com/hyuti/pocketbase-template/migrations"
	"github.com/hyuti/pocketbase-template/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase"
	"github.com/robfig/cron/v3"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// HTTP Server
	handler := controller.NewHandler(cfg)

	// Register cmd
	cmd.RegisterCMD(handler, l, cfg)

	controller.RegisterRoutes(handler, l)

	// Web api
	txtGenerator := webapi.NewDummyTextGenerator()

	// BackgroundTask
	taskClient := task.NewClient(cfg.Redis.URL, cfg.Redis.Password, cfg.Redis.DB)
	defer taskClient.Close()
	workerHandler := task.NewHandler()
	// Register Tasks
	task.RegisterTask(workerHandler, l, txtGenerator)
	workerServer := task.NewServer(workerHandler, cfg.Redis.URL, cfg.Redis.Password, cfg.Redis.DB)

	// Register job
	sLer := schedule.NewScheduler(cfg)
	schedule.RegisterSchedule(
		sLer,
		task.GetSayHelloExecutor(taskClient),
	)

	// Register validation
	validation.RegisterValidation(handler, l, cfg)

	errFromBGServer := make(chan error, 1)
	errFromHTTPServer := make(chan error, 1)
	go startBackGroundTaskServer(workerServer, errFromBGServer)
	go startHTTPServer(handler, errFromHTTPServer)
	go startJobSchedule(sLer)
	if err := <-errFromBGServer; err != nil {
		l.Fatal(fmt.Errorf("internal.app.Run: %w", err))
	}
	if err := <-errFromHTTPServer; err != nil {
		l.Fatal(fmt.Errorf("internal.app.Run: %w", err))
	}

}

func startBackGroundTaskServer(s *task.Server, errChan chan error) {
	errChan <- s.Run()
}

func startHTTPServer(s *pocketbase.PocketBase, errChan chan error) {
	errChan <- s.Start()
}

func startJobSchedule(handler *cron.Cron) {
	schedule.Start(handler)
}
