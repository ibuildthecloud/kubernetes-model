package model

const (
	STATUS_DETAILS_TYPE = "v1.StatusDetails"
)

type StatusDetails struct {
	Causes []StatusCause `json:"causes,omitempty" yaml:"causes,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RetryAfterSeconds int32 `json:"retryAfterSeconds,omitempty" yaml:"retry_after_seconds,omitempty"`
}
