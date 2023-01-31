package validation

import (
	"strings"
	"time"

	"github.com/hyuti/pocketbase-template/config"
	"github.com/hyuti/pocketbase-template/pkg/infrastructure/logger"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

func registerDemoBeforeCreate(hooks validatorMap) {
	type payload struct {
		PeriodMessage string `json:"period_message"`
	}

	registerValidator(hooks, "demo", func(ctx echo.Context, r *models.Record, i logger.Interface, c *config.Config) error {
		now := time.Now().Local().Hour()
		pl := new(payload)

		if err := ctx.Bind(pl); err != nil {
			return err
		}

		if now%2 == 0 {
			if strings.Compare(strings.ToLower(pl.PeriodMessage), "hello world") != 0 {
				return apis.NewBadRequestError("period message must be hello world", nil)
			}
		} else if strings.Compare(strings.ToLower(pl.PeriodMessage), "hello") != 0 {
			return apis.NewBadRequestError("period message must be hello", nil)
		}
		return nil
	})
}
func registerDemoBeforeUpdate(hooks validatorMap) {
	registerValidator(hooks, "demo", func(ctx echo.Context, r *models.Record, i logger.Interface, c *config.Config) error {
		return nil
	})
}
func registerDemoBeforeDelete(hooks validatorMap) {
	registerValidator(hooks, "demo", func(ctx echo.Context, r *models.Record, i logger.Interface, c *config.Config) error {
		return nil
	})
}
