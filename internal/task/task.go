package task

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/hyuti/pocketbase-template/internal/webapi"
	"github.com/hyuti/pocketbase-template/pkg/infrastructure/logger"
)

type (
	TaskHandler = func(context.Context, *asynq.Task) error
	Server      struct {
		server  *asynq.Server
		handler asynq.Handler
	}
)

func newTask[T any](pl T, typeTask string) (*asynq.Task, error) {
	payload, err := json.Marshal(pl)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(typeTask, payload), nil
}
func RegisterTask(handler *asynq.ServeMux, l logger.Interface, txtGenerator webapi.IDummyTextGenerator) {
	registerSayHelloHandler(handler, l, txtGenerator)
}

func NewHandler() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	return mux
}

func NewClient(url string, pwd string, db int) *asynq.Client {
	c := asynq.NewClient(
		asynq.RedisClientOpt{
			Addr:     url,
			Password: pwd,
			DB:       db,
		},
	)
	return c
}

func NewServer(handler asynq.Handler, url string, pwd string, db int) *Server {
	s := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     url,
			Password: pwd,
			DB:       db,
		},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	return &Server{
		server:  s,
		handler: handler,
	}
}

func (s *Server) Run() error {
	return s.server.Run(s.handler)
}
