// Package app configures and runs application.
package app

import (
	"github.com/hyuti/pocketbase-clean-template/cmd"
	"github.com/hyuti/pocketbase-clean-template/config"
	"github.com/hyuti/pocketbase-clean-template/internal/controller"
	_ "github.com/hyuti/pocketbase-clean-template/migrations"
	"github.com/hyuti/pocketbase-clean-template/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
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

	if err := handler.Start(); err != nil {
		l.Fatal(err)
	}
}
