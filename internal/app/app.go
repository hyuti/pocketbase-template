// Package app configures and runs application.
package app

import (
	"fmt"

	"github.com/hyuti/pocketbase-template/cmd"
	"github.com/hyuti/pocketbase-template/config"
	"github.com/hyuti/pocketbase-template/internal/controller"
	"github.com/hyuti/pocketbase-template/internal/schedule"
	"github.com/hyuti/pocketbase-template/internal/task"
	"github.com/hyuti/pocketbase-template/internal/webapi"
	_ "github.com/hyuti/pocketbase-template/migrations"
	"github.com/hyuti/pocketbase-template/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/robfig/cron/v3"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// HTTP Server
	handler := controller.NewHandler(cfg)

	// Register cmd
	cmd.RegisterCMD(handler, l, cfg)

	migratecmd.MustRegister(handler, handler.RootCmd, &migratecmd.Options{
		Automigrate: true,
	})

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

	go startBackGroundTaskServer(workerServer, l)
	startJobSchedule(sLer)
	if err := handler.Start(); err != nil {
		l.Fatal(err)
	}
}
func startBackGroundTaskServer(s *task.Server, l logger.Interface) {
	if err := s.Run(); err != nil {
		l.Fatal(fmt.Errorf("app - Run - workerServer.Run: %w", err))
	}
}
func startJobSchedule(handler *cron.Cron) {
	schedule.Start(handler)
}
