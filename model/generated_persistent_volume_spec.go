package model

const (
	PERSISTENT_VOLUME_SPEC_TYPE = "v1.PersistentVolumeSpec"
)

type PersistentVolumeSpec struct {
	AccessModes []PersistentVolumeAccessMode `json:"accessModes,omitempty" yaml:"access_modes,omitempty"`

	AwsElasticBlockStore *AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty" yaml:"aws_elastic_block_store,omitempty"`

	Capacity map[string]interface{} `json:"capacity,omitempty" yaml:"capacity,omitempty"`

	ClaimRef *ObjectReference `json:"claimRef,omitempty" yaml:"claim_ref,omitempty"`

	GcePersistentDisk *GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty" yaml:"gce_persistent_disk,omitempty"`

	Glusterfs *GlusterfsVolumeSource `json:"glusterfs,omitempty" yaml:"glusterfs,omitempty"`

	HostPath *HostPathVolumeSource `json:"hostPath,omitempty" yaml:"host_path,omitempty"`

	Iscsi *ISCSIVolumeSource `json:"iscsi,omitempty" yaml:"iscsi,omitempty"`

	Nfs *NFSVolumeSource `json:"nfs,omitempty" yaml:"nfs,omitempty"`

	PersistentVolumeReclaimPolicy string `json:"persistentVolumeReclaimPolicy,omitempty" yaml:"persistent_volume_reclaim_policy,omitempty"`

	Rbd *RBDVolumeSource `json:"rbd,omitempty" yaml:"rbd,omitempty"`
}
