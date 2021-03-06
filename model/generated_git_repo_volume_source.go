package model

const (
	GIT_REPO_VOLUME_SOURCE_TYPE = "v1.GitRepoVolumeSource"
)

type GitRepoVolumeSource struct {
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	Revision string `json:"revision,omitempty" yaml:"revision,omitempty"`
}
