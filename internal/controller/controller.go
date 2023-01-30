package controller

import (
	"net/http"

	"github.com/hyuti/pocketbase-clean-template/config"
	"github.com/hyuti/pocketbase-clean-template/pkg/infrastructure/logger"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func NewHandler(cfg *config.Config) *pocketbase.PocketBase {
	handler := pocketbase.NewWithConfig(&pocketbase.Config{
		DefaultDebug:   cfg.Debug,
		DefaultDataDir: cfg.DataDir,
	})
	return handler
}
func RegisterRoutes(handler core.App, l logger.Interface) {
	handler.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		router := Router{
			handler: e.Router,
		}
		RegisterHealthRoute(&router)
		return nil
	})
}

func createRoute(method, path string, handler echo.HandlerFunc) *echo.Route {
	return &echo.Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}

type Router struct {
	handler *echo.Echo
}

func (s *Router) Get(path string, handler echo.HandlerFunc) {
	s.handler.AddRoute(createRoute(http.MethodGet, path, handler))
}
func (s *Router) Post(path string, handler echo.HandlerFunc) {
	s.handler.AddRoute(createRoute(http.MethodPost, path, handler))
}
func (s *Router) Put(path string, handler echo.HandlerFunc) {
	s.handler.AddRoute(createRoute(http.MethodPut, path, handler))
}
func (s *Router) Patch(path string, handler echo.HandlerFunc) {
	s.handler.AddRoute(createRoute(http.MethodPatch, path, handler))
}
func (s *Router) Delete(path string, handler echo.HandlerFunc) {
	s.handler.AddRoute(createRoute(http.MethodDelete, path, handler))
}
func (s *Router) Options(path string, handler echo.HandlerFunc) {
	s.handler.AddRoute(createRoute(http.MethodOptions, path, handler))
}
func (s *Router) Head(path string, handler echo.HandlerFunc) {
	s.handler.AddRoute(createRoute(http.MethodHead, path, handler))
}
