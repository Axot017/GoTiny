package port

type Config interface {
	BaseUrl() string

	LogJson() bool
}
