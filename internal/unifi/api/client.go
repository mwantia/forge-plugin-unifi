package api

type ConnectedClient struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	MACAddress     string `json:"macAddress"`
	IPAddress      string `json:"ipAddress"`
	ConnectedAt    string `json:"connectedAt,omitempty"`
	UplinkDeviceID string `json:"uplinkDeviceId,omitempty"`

	Access ConnectedClientAccess `json:"access"`
}

type ConnectedClientAccess struct {
	Type string `json:"type"`
}

type ConnectedClientDetails struct {
	ConnectedClient
}
