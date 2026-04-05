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
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"limit": map[string]any{
					"type":        "integer",
					"description": "Maximum number of local sites to return",
				},
				"offset": map[string]any{
					"type":        "integer",
					"description": "Number of local sites to skip for pagination",
				},
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
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"site_id": map[string]any{
					"type":        "string",
					"description": "The site ID the adopted devices belongs to",
				},
				"limit": map[string]any{
					"type":        "integer",
					"description": "Maximum number of adopted devices to return",
				},
				"offset": map[string]any{
					"type":        "integer",
					"description": "Number of adopted devices to skip for pagination",
				},
			},
			"required": []any{"site_id"},
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
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"site_id": map[string]any{
					"type":        "string",
					"description": "The site ID the adopted device belongs to",
				},
				"device_id": map[string]any{
					"type":        "string",
					"description": "The adopted device ID to retrieve",
				},
			},
			"required": []any{"site_id", "device_id"},
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
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"site_id": map[string]any{
					"type":        "string",
					"description": "The site ID to query connected clients for",
				},
				"limit": map[string]any{
					"type":        "integer",
					"description": "Maximum number of connected clients to return per page",
				},
				"offset": map[string]any{
					"type":        "integer",
					"description": "Number of connected clients to skip for pagination",
				},
			},
			"required": []any{"site_id"},
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
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"site_id": map[string]any{
					"type":        "string",
					"description": "The site ID to connected client belongs to",
				},
				"client_id": map[string]any{
					"type":        "string",
					"description": "The connected client ID to retrieve",
				},
			},
			"required": []any{"site_id", "client_id"},
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
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"site_id": map[string]any{
					"type":        "string",
					"description": "The site ID to query networks for",
				},
				"limit": map[string]any{
					"type":        "integer",
					"description": "Maximum number of networks to return per page",
				},
				"offset": map[string]any{
					"type":        "integer",
					"description": "Number of networks to skip for pagination",
				},
			},
			"required": []any{"site_id"},
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
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"site_id": map[string]any{
					"type":        "string",
					"description": "The site ID the network belongs to",
				},
				"network_id": map[string]any{
					"type":        "string",
					"description": "The network ID to retrieve",
				},
			},
			"required": []any{"site_id", "network_id"},
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
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"site_id": map[string]any{
					"type":        "string",
					"description": "The site ID to query wifi broadcasts for",
				},
				"limit": map[string]any{
					"type":        "integer",
					"description": "Maximum number of wifi broadcasts to return per page",
				},
				"offset": map[string]any{
					"type":        "integer",
					"description": "Number of wifi broadcasts to skip for pagination",
				},
			},
			"required": []any{"site_id"},
		},
	},
	// /v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}
	"get_wifi_broadcast_details": {
		Name:        "get_wifi_broadcast_details",
		Description: "Retrieve detailed information about a specific Wifi, including name, network ID, security information (except passphrase), isolation configuration and if the wifi broadcast is enabled or not",
		Tags:        []string{"unifi", "network", "wifi", "broadcast", "details"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentFrequent,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"site_id": map[string]any{
					"type":        "string",
					"description": "The site ID the wifi broadcast belongs to",
				},
				"wifi_broadcast_id": map[string]any{
					"type":        "string",
					"description": "The wifi broadcast ID to retrieve",
				},
			},
			"required": []any{"site_id", "wifi_broadcast_id"},
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
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"site_id": map[string]any{
					"type":        "string",
					"description": "The site ID to query firewall zones for",
				},
				"limit": map[string]any{
					"type":        "integer",
					"description": "Maximum number of firewall zones to return per page",
				},
				"offset": map[string]any{
					"type":        "integer",
					"description": "Number of firewall zones to skip for pagination",
				},
			},
			"required": []any{"site_id"},
		},
	},
	// /v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}
	"get_firewall_zone": {
		Name:        "get_firewall_zone",
		Description: "Retrieve detailed information about a specific filewall zone, including name, the network ids associated with this firewall zone and metadata information - This does NOT provide more detailed informations than 'list_firewall_zones' already does",
		Tags:        []string{"unifi", "network", "firewall", "zone", "details"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentFrequent,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"site_id": map[string]any{
					"type":        "string",
					"description": "The site ID the wifi broadcast belongs to",
				},
				"firewall_zone_id": map[string]any{
					"type":        "string",
					"description": "The wifi broadcast ID to retrieve",
				},
			},
			"required": []any{"site_id", "firewall_zone_id"},
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
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"site_id": map[string]any{
					"type":        "string",
					"description": "The site ID to query firewall policies for",
				},
				"limit": map[string]any{
					"type":        "integer",
					"description": "Maximum number of firewall policies to return per page",
				},
				"offset": map[string]any{
					"type":        "integer",
					"description": "Number of firewall policies to skip for pagination",
				},
			},
			"required": []any{"site_id"},
		},
	},
	// /v1/sites/{siteId}/wifi/broadcasts/{wifiBroadcastId}
	"get_firewall_policy": {
		Name:        "get_firewall_policy",
		Description: "Retrieve information about a specific filewall policy, including name, the network ids associated with this firewall policy and metadata information - This does NOT provide more detailed informations than 'list_firewall_policies' already does",
		Tags:        []string{"unifi", "network", "firewall", "policy"},
		Annotations: plugins.ToolAnnotations{
			ReadOnly:              true,
			Idempotent:            false,
			IdempotentProbability: plugins.ToolIdempotentFrequent,
			CostHint:              plugins.ToolCostCheap,
		},
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"site_id": map[string]any{
					"type":        "string",
					"description": "The site ID the firewall policy belongs to",
				},
				"firewall_policy_id": map[string]any{
					"type":        "string",
					"description": "The firewall policy ID to retrieve",
				},
			},
			"required": []any{"site_id", "firewall_policy_id"},
		},
	},
}
