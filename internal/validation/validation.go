package validation

import (
	"github.com/hyuti/pocketbase-template/config"
	"github.com/hyuti/pocketbase-template/pkg/infrastructure/logger"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

type (
	validator    func(echo.Context, *models.Record, logger.Interface, *config.Config) error
	validatorMap map[string]validator
)

func registerValidator(hooks validatorMap, n string, v validator) {
	hooks[n] = v
}

func registerBeforeCreateValidator(hooks validatorMap) {
	// add more before create validators here
	registerDemoBeforeCreate(hooks)
}

func registerBeforeUpdateValidator(hooks validatorMap) {
	// add more before update validators here
	registerDemoBeforeUpdate(hooks)
}

func registerBeforeDeleteValidator(hooks validatorMap) {
	// add more before delete validators here
	registerDemoBeforeDelete(hooks)
}

func registerBeforeCreate(handler core.App, l logger.Interface, cfg *config.Config) {
	hooks := make(validatorMap)

	registerBeforeCreateValidator(hooks)

	handler.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		clt := e.Record.Collection().Name
		validator, ok := hooks[clt]
		if ok {
			return validator(e.HttpContext, e.Record, l, cfg)
		}
		return nil
	})
}

func registerBeforeUpdate(handler core.App, l logger.Interface, cfg *config.Config) {
	hooks := make(validatorMap)

	registerBeforeUpdateValidator(hooks)

	handler.OnRecordBeforeUpdateRequest().Add(func(e *core.RecordUpdateEvent) error {
		clt := e.Record.Collection().Name
		validator, ok := hooks[clt]
		if ok {
			return validator(e.HttpContext, e.Record, l, cfg)
		}
		return nil
	})
}

func registerBeforeDelete(handler core.App, l logger.Interface, cfg *config.Config) {
	hooks := make(validatorMap)

	registerBeforeDeleteValidator(hooks)

	handler.OnRecordBeforeDeleteRequest().Add(func(e *core.RecordDeleteEvent) error {
		clt := e.Record.Collection().Name
		validator, ok := hooks[clt]
		if ok {
			return validator(e.HttpContext, e.Record, l, cfg)
		}
		return nil
	})
}

func RegisterValidation(handler core.App, l logger.Interface, cfg *config.Config) {
	registerBeforeCreate(handler, l, cfg)
	registerBeforeUpdate(handler, l, cfg)
	registerBeforeDelete(handler, l, cfg)
}
