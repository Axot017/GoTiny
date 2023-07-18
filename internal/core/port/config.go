package port

type Config interface {
	BaseUrl() string

	MaxTrackingDays() uint

	LogJson() bool
}
