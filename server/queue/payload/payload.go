package payload

type OnProvisionInstancePayload struct {
	FleetRequestId  string  `json:"fleetRequestId"`
	EipAllocationId *string `json:"eipAllocationId,omitempty"`
}