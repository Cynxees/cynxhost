package types

type StorageStatus string
type InstanceTypeStatus string
type InstanceStatus string
type PersistentNodeStatus string

const (
	StorageStatusInUseRunningInstance  StorageStatus = "IN_USE:RUNNING_INSTANCE"
	StorageStatusInUseCreatingSnapshot StorageStatus = "IN_USE:CREATING_SNAPSHOT"
	StorageStatusReady                 StorageStatus = "READY"
	StorageStatusNew                   StorageStatus = "NEW"
)
