package api

type WifiBroadcast struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Type    string `json:"type"`

	Network               WifiBroadcastNetwork               `json:"network"`
	SecurityConfiguration WifiBroadcastSecurityConfiguration `json:"securityConfiguration"`
	Metadata              map[string]any                     `json:"metadata"`
}

type WifiBroadcastNetwork struct {
	Type      string `json:"type"`
	NetworkID string `json:"networkId"`
}

type WifiBroadcastSecurityConfiguration struct {
	Type string `json:"type"`
}

type WifiBroadcastDetails struct {
	ID                                  string `json:"id"`
	Name                                string `json:"name"`
	Enabled                             bool   `json:"enabled"`
	Type                                string `json:"type"`
	MulticastToUnicastConversionEnabled bool   `json:"multicastToUnicastConversionEnabled"`
	ClientIsolationEnabled              bool   `json:"clientIsolationEnabled"`
	HideName                            bool   `json:"hideName"`
	UapsdEnabled                        bool   `json:"uapsdEnabled"`
	BandSteeringEnabled                 bool   `json:"bandSteeringEnabled"`
	ArpProxyEnabled                     bool   `json:"arpProxyEnabled"`
	BssTransitionEnabled                bool   `json:"bssTransitionEnabled"`
	AdvertiseDeviceName                 bool   `json:"advertiseDeviceName"`

	Network               WifiBroadcastDetailsNetwork               `json:"network"`
	SecurityConfiguration WifiBroadcastDetailsSecurityConfiguration `json:"securityConfiguration"`
	Metadata              map[string]any                            `json:"metadata"`
}

type WifiBroadcastDetailsNetwork struct {
	Type      string `json:"type"`
	NetworkID string `json:"networkId"`
}

type WifiBroadcastDetailsSecurityConfiguration struct {
	Type                      string `json:"type"`
	GroupRekeyIntervalSeconds int    `json:"groupRekeyIntervalSeconds"`
	FastRoamingEnabled        bool   `json:"fastRoamingEnabled"`
	PmfMode                   string `json:"pmfMode"`
	Wpa3FastRoamingEnabled    bool   `json:"wpa3FastRoamingEnabled"`
}
