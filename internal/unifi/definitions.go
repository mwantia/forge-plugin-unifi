package unifi

import "github.com/mwantia/forge-sdk/pkg/plugins"

var toolDefinitions = map[string]plugins.ToolDefinition{
	// /v1/sites
	"list_local_sites": {
		Name:        "list_local_sites",
		Description: "Retrieve a paginated list of local sites managed by this Network application. Site ID is required for other UniFi Network API calls",
		Tags:        []string{"unifi", "network", "sites"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            true,
			IdempotentProbability: plugins.ToolIdempotentGuaranteed,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: plugins.ToolParameters{
			Type: "object",
			Properties: map[string]plugins.ToolProperty{
				"limit":  {Type: "integer", Description: "Maximum number of local sites to return"},
				"offset": {Type: "integer", Description: "Number of local sites to skip for pagination"},
			},
		},
	},

	// /v1/sites/{siteId}/devices
	"list_adopted_devices": {
		Name:        "list_adopted_devices",
		Description: "Retrieve a paginated list of all adopted devices on a site, including basic device information",
		Tags:        []string{"unifi", "network", "devices", "list"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentFrequent,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: plugins.ToolParameters{
			Type: "object",
			Properties: map[string]plugins.ToolProperty{
				"site_id": {Type: "string", Description: "The site ID the adopted devices belongs to", Format: "123e4567-e89b-12d3-a456-426614174000"},
				"limit":   {Type: "integer", Description: "Maximum number of adopted devices to return"},
				"offset":  {Type: "integer", Description: "Number of adopted devices to skip for pagination"},
			},
			Required: []string{"site_id"},
		},
	},

	// /v1/sites/{siteId}/devices/{deviceId}
	"get_adopted_device_details": {
		Name:        "get_adopted_device_details",
		Description: "Retrieve detailed information about a specific adopted device, including firmware versioning, uplink state, details about device features and interfaces (ports, radios) and other key attributes",
		Tags:        []string{"unifi", "network", "devices", "details"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentFrequent,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: plugins.ToolParameters{
			Type: "object",
			Properties: map[string]plugins.ToolProperty{
				"site_id":   {Type: "string", Description: "The site ID the adopted device belongs to", Format: "123e4567-e89b-12d3-a456-426614174000"},
				"device_id": {Type: "string", Description: "The adopted device ID to retrieve", Format: "123e4567-e89b-12d3-a456-426614174000"},
			},
			Required: []string{"site_id", "device_id"},
		},
	},

	// /v1/sites/{siteId}/clients
	"list_connected_clients": {
		Name:        "list_connected_clients",
		Description: "Retrieve a paginated list of all connected clients on a site, including physical devices (computers, smartphones) and active VPN connections",
		Tags:        []string{"unifi", "network", "clients", "list"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentUnlikely,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: plugins.ToolParameters{
			Type: "object",
			Properties: map[string]plugins.ToolProperty{
				"site_id": {Type: "string", Description: "The site ID to query connected clients for", Format: "123e4567-e89b-12d3-a456-426614174000"},
				"limit":   {Type: "integer", Description: "Maximum number of connected clients to return per page"},
				"offset":  {Type: "integer", Description: "Number of connected clients to skip for pagination"},
			},
			Required: []string{"site_id"},
		},
	},

	// /v1/sites/{siteId}/clients/{clientId}
	"get_connected_client_details": {
		Name:        "get_connected_client_details",
		Description: "Retrieve detailed information about a specific connected client, including name, IP address, MAC address, connection type and access information",
		Tags:        []string{"unifi", "network", "clients", "details"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentUnlikely,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: plugins.ToolParameters{
			Type: "object",
			Properties: map[string]plugins.ToolProperty{
				"site_id":   {Type: "string", Description: "The site ID the connected client belongs to", Format: "123e4567-e89b-12d3-a456-426614174000"},
				"client_id": {Type: "string", Description: "The connected client ID to retrieve", Format: "123e4567-e89b-12d3-a456-426614174000"},
			},
			Required: []string{"site_id", "client_id"},
		},
	},

	// /v1/sites/{siteId}/networks
	"list_networks": {
		Name:        "list_networks",
		Description: "Retrieve a paginated list of all Networks on a site",
		Tags:        []string{"unifi", "network", "networks", "list"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentFrequent,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: plugins.ToolParameters{
			Type: "object",
			Properties: map[string]plugins.ToolProperty{
				"site_id": {Type: "string", Description: "The site ID to query networks for", Format: "123e4567-e89b-12d3-a456-426614174000"},
				"limit":   {Type: "integer", Description: "Maximum number of networks to return per page"},
				"offset":  {Type: "integer", Description: "Number of networks to skip for pagination"},
			},
			Required: []string{"site_id"},
		},
	},

	// /v1/sites/{siteId}/networks/{networkId}
	"get_network_details": {
		Name:        "get_network_details",
		Description: "Retrieve detailed information about a specific network, including name, vlan ID, metadata information, ipv4- and ipv6-configuration and if the network is enabled or not",
		Tags:        []string{"unifi", "network", "networks", "details"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentUnlikely,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: plugins.ToolParameters{
			Type: "object",
			Properties: map[string]plugins.ToolProperty{
				"site_id":    {Type: "string", Description: "The site ID the network belongs to", Format: "123e4567-e89b-12d3-a456-426614174000"},
				"network_id": {Type: "string", Description: "The network ID to retrieve", Format: "123e4567-e89b-12d3-a456-426614174000"},
			},
			Required: []string{"site_id", "network_id"},
		},
	},

	// /v1/sites/{siteId}/wifi/broadcasts
	"list_wifi_broadcast": {
		Name:        "list_wifi_broadcast",
		Description: "Retrieve a paginated list of all Wifi Broadcasts on a site",
		Tags:        []string{"unifi", "network", "wifi", "broadcast", "list"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentFrequent,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: plugins.ToolParameters{
			Type: "object",
			Properties: map[string]plugins.ToolProperty{
				"site_id": {Type: "string", Description: "The site ID to query wifi broadcasts for", Format: "123e4567-e89b-12d3-a456-426614174000"},
				"limit":   {Type: "integer", Description: "Maximum number of wifi broadcasts to return per page"},
				"offset":  {Type: "integer", Description: "Number of wifi broadcasts to skip for pagination"},
			},
			Required: []string{"site_id"},
		},
	},

	// /v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}
	"get_wifi_broadcast_details": {
		Name:        "get_wifi_broadcast_details",
		Description: "Retrieve detailed information about a specific Wifi broadcast, including name, network ID, security information (except passphrase), isolation configuration and if the wifi broadcast is enabled or not",
		Tags:        []string{"unifi", "network", "wifi", "broadcast", "details"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentFrequent,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: plugins.ToolParameters{
			Type: "object",
			Properties: map[string]plugins.ToolProperty{
				"site_id":           {Type: "string", Description: "The site ID the wifi broadcast belongs to", Format: "123e4567-e89b-12d3-a456-426614174000"},
				"wifi_broadcast_id": {Type: "string", Description: "The wifi broadcast ID to retrieve", Format: "123e4567-e89b-12d3-a456-426614174000"},
			},
			Required: []string{"site_id", "wifi_broadcast_id"},
		},
	},

	// /v1/sites/{siteId}/firewall/zones
	"list_firewall_zones": {
		Name:        "list_firewall_zones",
		Description: "Retrieve a list of all firewall zones on a site",
		Tags:        []string{"unifi", "network", "firewall", "zone", "list"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentFrequent,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: plugins.ToolParameters{
			Type: "object",
			Properties: map[string]plugins.ToolProperty{
				"site_id": {Type: "string", Description: "The site ID to query firewall zones for", Format: "123e4567-e89b-12d3-a456-426614174000"},
				"limit":   {Type: "integer", Description: "Maximum number of firewall zones to return per page"},
				"offset":  {Type: "integer", Description: "Number of firewall zones to skip for pagination"},
			},
			Required: []string{"site_id"},
		},
	},

	// /v1/sites/{siteId}/firewall/zones/{firewallZoneId}
	"get_firewall_zone": {
		Name:        "get_firewall_zone",
		Description: "Retrieve detailed information about a specific firewall zone, including name and the network IDs associated with it",
		Tags:        []string{"unifi", "network", "firewall", "zone", "details"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentFrequent,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: plugins.ToolParameters{
			Type: "object",
			Properties: map[string]plugins.ToolProperty{
				"site_id":          {Type: "string", Description: "The site ID the firewall zone belongs to", Format: "123e4567-e89b-12d3-a456-426614174000"},
				"firewall_zone_id": {Type: "string", Description: "The firewall zone ID to retrieve", Format: "123e4567-e89b-12d3-a456-426614174000"},
			},
			Required: []string{"site_id", "firewall_zone_id"},
		},
	},

	// /v1/sites/{siteId}/firewall/policies
	"list_firewall_policies": {
		Name:        "list_firewall_policies",
		Description: "Retrieve a list of all firewall policies on a site",
		Tags:        []string{"unifi", "network", "firewall", "policies", "list"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentFrequent,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: plugins.ToolParameters{
			Type: "object",
			Properties: map[string]plugins.ToolProperty{
				"site_id": {Type: "string", Description: "The site ID to query firewall policies for", Format: "123e4567-e89b-12d3-a456-426614174000"},
				"limit":   {Type: "integer", Description: "Maximum number of firewall policies to return per page"},
				"offset":  {Type: "integer", Description: "Number of firewall policies to skip for pagination"},
			},
			Required: []string{"site_id"},
		},
	},

	// /v1/sites/{siteId}/firewall/policies/{firewallPolicyId}
	"get_firewall_policy": {
		Name:        "get_firewall_policy",
		Description: "Retrieve detailed information about a specific firewall policy, including name and the network IDs associated with it",
		Tags:        []string{"unifi", "network", "firewall", "policy"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentFrequent,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: plugins.ToolParameters{
			Type: "object",
			Properties: map[string]plugins.ToolProperty{
				"site_id":            {Type: "string", Description: "The site ID the firewall policy belongs to", Format: "123e4567-e89b-12d3-a456-426614174000"},
				"firewall_policy_id": {Type: "string", Description: "The firewall policy ID to retrieve", Format: "123e4567-e89b-12d3-a456-426614174000"},
			},
			Required: []string{"site_id", "firewall_policy_id"},
		},
	},
}
