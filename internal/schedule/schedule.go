package schedule

import (
	"github.com/hyuti/pocketbase-template/config"
	"github.com/hyuti/pocketbase-template/internal/task"
	"github.com/robfig/cron/v3"
)

const (
	everyMinute = "@every 60s"
)

func NewScheduler(cfg *config.Config) *cron.Cron {
	c := cron.New()
	return c
}

func Start(handler *cron.Cron) {
	handler.Start()
}

func registerJob(handler *cron.Cron, spec string, cmd func()) {
	handler.AddFunc(spec, cmd)
}

func RegisterSchedule(handler *cron.Cron, taskXor task.IExecuteTask[*task.SayHelloPayload]) {
	registerSayHelloJob(handler, taskXor)
}
