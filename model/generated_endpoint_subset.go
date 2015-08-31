package model

const (
	ENDPOINT_SUBSET_TYPE = "v1.EndpointSubset"
)

type EndpointSubset struct {
	Addresses []EndpointAddress `json:"addresses,omitempty" yaml:"addresses,omitempty"`

	Ports []EndpointPort `json:"ports,omitempty" yaml:"ports,omitempty"`
}
