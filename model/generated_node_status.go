package model

const (
	NODE_STATUS_TYPE = "v1.NodeStatus"
)

type NodeStatus struct {
	Addresses []NodeAddress `json:"addresses,omitempty" yaml:"addresses,omitempty"`

	Capacity map[string]interface{} `json:"capacity,omitempty" yaml:"capacity,omitempty"`

	Conditions []NodeCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`

	NodeInfo *NodeSystemInfo `json:"nodeInfo,omitempty" yaml:"node_info,omitempty"`

	Phase string `json:"phase,omitempty" yaml:"phase,omitempty"`
}
