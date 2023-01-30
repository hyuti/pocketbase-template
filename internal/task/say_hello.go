package task

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/hyuti/pocketbase-template/internal/webapi"
	"github.com/hyuti/pocketbase-template/pkg/infrastructure/logger"
)

const (
	typeSayHello = "notify:sayHello"
)

type (
	SayHelloExecutor struct {
		client *asynq.Client
	}
	SayHelloPayload struct {
		Message string
	}
)

func (s *SayHelloExecutor) ExecuteTask(pl *SayHelloPayload) error {
	if pl == nil {
		pl = new(SayHelloPayload)
	}
	task, err := newTask(pl, typeSayHello)
	if err != nil {
		return err
	}
	_, err = s.client.Enqueue(task)
	if err != nil {
		return err
	}
	return nil
}
func GetSayHelloExecutor(c *asynq.Client) IExecuteTask[*SayHelloPayload] {
	return &SayHelloExecutor{
		client: c,
	}
}
func registerSayHelloHandler(handler *asynq.ServeMux, l logger.Interface, txtGenerator webapi.IDummyTextGenerator) {
	handler.HandleFunc(typeSayHello, sayHelloTaskHandlerWrapper(l, txtGenerator))
}

func sayHelloTaskHandlerWrapper(l logger.Interface, txtGenerator webapi.IDummyTextGenerator) TaskHandler {
	return func(ctx context.Context, t *asynq.Task) error {
		p := new(SayHelloPayload)
		if err := json.Unmarshal(t.Payload(), p); err != nil {
			return err
		}
		dummyTxt, err := txtGenerator.Get()
		if err != nil {
			return err
		}
		l.Info(fmt.Sprintf("%s: %s", p.Message, dummyTxt))
		return nil
	}
}
