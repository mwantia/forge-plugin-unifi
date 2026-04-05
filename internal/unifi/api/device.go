package api

type AdoptedDevice struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Model             string   `json:"model,omitempty"`
	MACAddress        string   `json:"macAddress"`
	IPAddress         string   `json:"ipAddress"`
	State             string   `json:"state,omitempty"`
	Supported         bool     `json:"supported,omitempty"`
	FirmwareVersion   string   `json:"firmwareVersion,omitempty"`
	FirmwareUpdatable bool     `json:"firmwareUpdatable,omitempty"`
	Features          []string `json:"features,omitempty"`
	Interfaces        []string `json:"interfaces,omitempty"`
}

type AdoptedDeviceDetails struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Model             string `json:"model,omitempty"`
	MACAddress        string `json:"macAddress"`
	IPAddress         string `json:"ipAddress"`
	State             string `json:"state,omitempty"`
	Supported         bool   `json:"supported,omitempty"`
	FirmwareVersion   string `json:"firmwareVersion,omitempty"`
	FirmwareUpdatable bool   `json:"firmwareUpdatable,omitempty"`
	ProvisionedAt     string `json:"provisionedAt,omitempty"`

	Uplink     AdoptedDeviceDetailsUplink     `json:"uplink"`
	Interfaces AdoptedDeviceDetailsInterfaces `json:"interfaces"`
}

type AdoptedDeviceDetailsUplink struct {
	DeviceID string `json:"deviceId"`
}

type AdoptedDeviceDetailsInterfaces struct {
	Ports []AdoptedDeviceDetailsInterfacePort `json:"ports"`
}

type AdoptedDeviceDetailsInterfacePort struct {
	IDX          int    `json:"idx"`
	State        string `json:"state,omitempty"`
	Connector    string `json:"connector,omitempty"`
	SpeedMbps    int    `json:"speedMbps"`
	MaxSpeedMbps int    `json:"maxSpeedMbps"`

	PoE AdoptedDeviceDetailsInterfacePortPoE `json:"poe"`
}

type AdoptedDeviceDetailsInterfacePortPoE struct {
	Enabled  bool   `json:"enabled"`
	State    string `json:"state"`
	Type     int    `json:"type"`
	Standard string `json:"standard"`
}
