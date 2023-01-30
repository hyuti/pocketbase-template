package schedule

import (
	"fmt"

	"github.com/hyuti/pocketbase-template/internal/task"
	"github.com/robfig/cron/v3"
)

func registerSayHelloJob(handler *cron.Cron, taskXor task.IExecuteTask[*task.SayHelloPayload]) {
	generationTimes := 0
	registerJob(handler, everyMinute, func() {
		taskXor.ExecuteTask(&task.SayHelloPayload{
			Message: fmt.Sprintf("%d Generation", generationTimes),
		})
		generationTimes += 1
	})
}
