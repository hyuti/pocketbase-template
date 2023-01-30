package task

type (
	IExecuteTask[T any] interface {
		ExecuteTask(T) error
	}
)
