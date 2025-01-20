package types

type LaunchCallbackPersistentNodeType string

const (
	LaunchCallbackPersistentNodeTypeInitialLaunch LaunchCallbackPersistentNodeType = "INITIAL_LAUNCH"
)

type StatusCallbackPersistentNodeType string

const (
	SetupSuccessCallbackPersistentNodeType StatusCallbackPersistentNodeType = "SETUP_SUCCESS"
)

type ScriptVariableType string

const (
	ScriptVariableTypeOption ScriptVariableType = "OPTION"
)
