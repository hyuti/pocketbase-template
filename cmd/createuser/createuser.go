package createuser

import (
	"fmt"

	"github.com/hyuti/pocketbase-clean-template/config"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/spf13/cobra"
)

func getOrCreateAdmin(app core.App, email *string) (*models.Admin, error) {
	ins, _ := app.Dao().FindAdminByEmail(*email)
	return ins, nil
}

func CreateUser(app core.App, cfg *config.Config) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		admin, err := getOrCreateAdmin(app, &cfg.Email)
		if err != nil {
			fmt.Printf("error creating admin: %s\n", err)
		}
		if admin == nil {
			admin = new(models.Admin)
			form := forms.NewAdminUpsert(app, admin)
			form.Email = cfg.Email
			form.Password = cfg.Password
			form.PasswordConfirm = cfg.Password

			err = form.Submit()

			if err != nil {
				fmt.Printf("error creating admin: %s\n", err)
			}
		} else {
			fmt.Printf("skipping creating admin...")
		}
	}
}
