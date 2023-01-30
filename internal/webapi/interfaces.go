package webapi

type (
	IDummyTextGenerator interface {
		Get() (string, error)
	}
)
