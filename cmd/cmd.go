package cmd

import (
	"github.com/hyuti/pocketbase-template/cmd/createuser"
	"github.com/hyuti/pocketbase-template/config"
	"github.com/hyuti/pocketbase-template/pkg/infrastructure/logger"
	"github.com/pocketbase/pocketbase"
	"github.com/spf13/cobra"
)

func RegisterCMD(
	handler *pocketbase.PocketBase,
	l logger.Interface,
	cfg *config.Config,
) {
	handler.RootCmd.AddCommand(&cobra.Command{
		Use: "createuser",
		Run: createuser.CreateUser(handler, cfg),
	})

}
