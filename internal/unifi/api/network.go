package api

type Network struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Enabled    bool           `json:"enabled"`
	Default    bool           `json:"default"`
	Management string         `json:"management"`
	VlanID     int            `json:"vlanId"`
	ZoneID     string         `json:"zoneId"`
	Metadata   map[string]any `json:"metadata"`
}

type NetworkDetails struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	Enabled               bool   `json:"enabled"`
	Default               bool   `json:"default"`
	Management            string `json:"management"`
	VlanID                int    `json:"vlanId"`
	ZoneID                string `json:"zoneId"`
	IsolationEnabled      bool   `json:"isolationEnabled"`
	CellularBackupEnabled bool   `json:"cellularBackupEnabled"`
	InternetAccessEnabled bool   `json:"internetAccessEnabled"`
	MdnsForwardingEnabled bool   `json:"mdnsForwardingEnabled"`

	IPv4Configuration NetworkDetailsIPv4Configuration `json:"ipv4Configuration"`
	IPv6Configuration NetworkDetailsIPv6Configuration `json:"ipv6Configuration"`
	Metadata          map[string]any                  `json:"metadata"`
}

type NetworkDetailsIPv4Configuration struct {
	HostIpAddress    string `json:"hostIpAddress"`
	PrefixLength     int    `json:"prefixLength"`
	AutoScaleEnabled bool   `json:"autoScaleEnabled"`

	DhcpConfiguration NetworkDetailsIPv4DhcpConfiguration `json:"dhcpConfiguration"`
}

type NetworkDetailsIPv4DhcpConfiguration struct {
	Mode                         string   `json:"mode"`
	DnsServerIpAddressesOverride []string `json:"dnsServerIpAddressesOverride"`
	LeaseTimeSeconds             int      `json:"leaseTimeSeconds"`
	DomainName                   string   `json:"domainName"`
	PingConflictDetectionEnabled bool     `json:"pingConflictDetectionEnabled"`
	Option43Value                string   `json:"option43Value"`

	IpAddressRange NetworkDetailsIPv4DhcpIpAddressRangeConfiguration `json:"ipAddressRange"`
}

type NetworkDetailsIPv4DhcpIpAddressRangeConfiguration struct {
	Start string `json:"start"`
	Stop  string `json:"stop"`
}

type NetworkDetailsIPv6Configuration struct {
	InterfaceType                string   `json:"interfaceType"`
	HostIpAddress                string   `json:"hostIpAddress"`
	PrefixLength                 int      `json:"prefixLength"`
	DnsServerIpAddressesOverride []string `json:"dnsServerIpAddressesOverride"`
}
