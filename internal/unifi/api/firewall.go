package api

type FirewallZone struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	NetworkIDs []string       `json:"networkIds"`
	Metadata   map[string]any `json:"metadata"`
}

type FirewallPolicy struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Enabled        bool   `json:"enabled"`
	Index          int    `json:"index"`
	LoggingEnabled bool   `json:"loggingEnabled"`

	Action          FirewallPolicyAction          `json:"action"`
	Source          FirewallPolicySource          `json:"source"`
	Destination     FirewallPolicyDestination     `json:"destination"`
	IpProtocolScope FirewallPolicyIpProtocolScope `json:"ipProtocolScope"`
	Metadata        map[string]any                `json:"metadata"`
}

type FirewallPolicyAction struct {
	Type               string `json:"type"`
	AllowReturnTraffic bool   `json:"allowReturnTraffic"`
}

type FirewallPolicySource struct {
	ZoneID string `json:"zoneId"`
}

type FirewallPolicyDestination struct {
	ZoneID string `json:"zoneId"`
}

type FirewallPolicyIpProtocolScope struct {
	IpVersion string `json:"ipVersion"`
}
