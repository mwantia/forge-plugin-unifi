package unifi

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/mwantia/forge-plugin-unifi/internal/unifi/api"
	"github.com/mwantia/forge-sdk/pkg/plugins"
)

type UnifiToolPlugin struct {
	plugins.UnimplementedToolsPlugin

	driver *UnifiDriver
	client *http.Client
}

type apiResponse[T any] struct {
	Data       []T `json:"data"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	Count      int `json:"count"`
	TotalCount int `json:"totalCount"`
}

func NewUnifiToolPlugin(driver *UnifiDriver) (*UnifiToolPlugin, error) {
	timeout, err := time.ParseDuration(driver.config.Timeout)
	if err != nil {
		return nil, fmt.Errorf("invalid timeout %q: %w", driver.config.Timeout, err)
	}

	return &UnifiToolPlugin{
		driver: driver,
		client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: driver.config.InsecureSkipTLS,
				},
			},
		},
	}, nil
}

func (p *UnifiToolPlugin) GetLifecycle() plugins.Lifecycle {
	return p.driver
}

func (p *UnifiToolPlugin) ListTools(_ context.Context, filter plugins.ListToolsFilter) (*plugins.ListToolsResponse, error) {
	tools := make([]plugins.ToolDefinition, 0)
	for _, def := range toolDefinitions {
		if plugins.MatchesToolsFilter(def, filter) {
			tools = append(tools, def)
		}
	}

	return &plugins.ListToolsResponse{
		Tools: tools,
	}, nil
}

func (p *UnifiToolPlugin) GetTool(_ context.Context, name string) (*plugins.ToolDefinition, error) {
	def, ok := toolDefinitions[strings.ToLower(name)]
	if !ok {
		return nil, fmt.Errorf("tool %q not found", name)
	}

	return &def, nil
}

func (p *UnifiToolPlugin) Validate(_ context.Context, req plugins.ExecuteRequest) (*plugins.ValidateResponse, error) {
	def, ok := toolDefinitions[req.Tool]
	if !ok {
		return &plugins.ValidateResponse{
			Valid:  false,
			Errors: []string{fmt.Sprintf("unknown tool %q", req.Tool)},
		}, nil
	}
	return plugins.ValidateAgainstDefinition(def, req), nil
}

func (p *UnifiToolPlugin) Execute(ctx context.Context, req plugins.ExecuteRequest) (*plugins.ExecuteResponse, error) {
	switch req.Tool {
	case "list_local_sites":
		return p.execList(ctx, "/sites", req.Arguments, func() any { return &apiResponse[api.LocalSite]{} },
			func(r any) any {
				res := r.(*apiResponse[api.LocalSite])
				return PaginatedResult("sites", res.Data, res)
			})

	case "list_adopted_devices":
		return p.execListForSite(ctx, "devices", req.Arguments, func() any { return &apiResponse[api.AdoptedDevice]{} },
			func(r any) any {
				res := r.(*apiResponse[api.AdoptedDevice])
				return PaginatedResult("devices", res.Data, res)
			})

	case "get_adopted_device_details":
		return p.execGetForSite(ctx, "devices", "device_id", req.Arguments, &api.AdoptedDeviceDetails{})

	case "list_connected_clients":
		return p.execListForSite(ctx, "clients", req.Arguments, func() any { return &apiResponse[api.ConnectedClient]{} },
			func(r any) any {
				res := r.(*apiResponse[api.ConnectedClient])
				return PaginatedResult("clients", res.Data, res)
			})

	case "get_connected_client_details":
		return p.execGetForSite(ctx, "clients", "client_id", req.Arguments, &api.ConnectedClientDetails{})

	case "list_networks":
		return p.execListForSite(ctx, "networks", req.Arguments, func() any { return &apiResponse[api.Network]{} },
			func(r any) any {
				res := r.(*apiResponse[api.Network])
				return PaginatedResult("networks", res.Data, res)
			})

	case "get_network_details":
		return p.execGetForSite(ctx, "networks", "network_id", req.Arguments, &api.NetworkDetails{})

	case "list_wifi_broadcast":
		return p.execListForSite(ctx, "wifi/broadcasts", req.Arguments, func() any { return &apiResponse[api.WifiBroadcast]{} },
			func(r any) any {
				res := r.(*apiResponse[api.WifiBroadcast])
				return PaginatedResult("wifi_broadcasts", res.Data, res)
			})

	case "get_wifi_broadcast_details":
		return p.execGetForSite(ctx, "wifi/broadcasts", "wifi_broadcast_id", req.Arguments, &api.WifiBroadcastDetails{})

	case "list_firewall_zones":
		return p.execListForSite(ctx, "firewall/zones", req.Arguments, func() any { return &apiResponse[api.FirewallZone]{} },
			func(r any) any {
				res := r.(*apiResponse[api.FirewallZone])
				return PaginatedResult("firewall_zones", res.Data, res)
			})

	case "get_firewall_zone":
		return p.execGetForSite(ctx, "firewall/zones", "firewall_zone_id", req.Arguments, &api.FirewallZone{})

	case "list_firewall_policies":
		return p.execListForSite(ctx, "firewall/policies", req.Arguments, func() any { return &apiResponse[api.FirewallPolicy]{} },
			func(r any) any {
				res := r.(*apiResponse[api.FirewallPolicy])
				return PaginatedResult("firewall_policies", res.Data, res)
			})

	case "get_firewall_policy":
		return p.execGetForSite(ctx, "firewall/policies", "firewall_policy_id", req.Arguments, &api.FirewallPolicy{})

	default:
		return &plugins.ExecuteResponse{
			Result:  fmt.Sprintf("unknown tool: %s", req.Tool),
			IsError: true,
		}, nil
	}
}

// execList fetches a paginated list from path, decodes into newResult(), and shapes the output via shape().
func (p *UnifiToolPlugin) execList(ctx context.Context, path string, args map[string]any, newResult func() any, shape func(any) any) (*plugins.ExecuteResponse, error) {
	dst := newResult()
	if err := p.DoGetRequest(ctx, path+PaginationQuery(args), dst); err != nil {
		return &plugins.ExecuteResponse{
			Result:  err.Error(),
			IsError: true,
		}, nil
	}

	return &plugins.ExecuteResponse{
		Result: shape(dst),
	}, nil
}

// execListForSite fetches a paginated list scoped to a site.
func (p *UnifiToolPlugin) execListForSite(ctx context.Context, resource string, args map[string]any, newResult func() any, shape func(any) any) (*plugins.ExecuteResponse, error) {
	siteID, ok := args["site_id"].(string)
	if !ok || siteID == "" {
		return &plugins.ExecuteResponse{
			Result:  "site_id is required, but not provided",
			IsError: true,
		}, nil
	}

	return p.execList(ctx, fmt.Sprintf("/sites/%s/%s", siteID, resource), args, newResult, shape)
}

// execGetForSite fetches a single resource by ID, scoped to a site.
// dst must be a pointer to the target struct.
func (p *UnifiToolPlugin) execGetForSite(ctx context.Context, resource, idKey string, args map[string]any, dst any) (*plugins.ExecuteResponse, error) {
	siteID, ok := args["site_id"].(string)
	if !ok || siteID == "" {
		return &plugins.ExecuteResponse{
			Result:  "site_id is required, but not provided",
			IsError: true,
		}, nil
	}
	resourceID, ok := args[idKey].(string)
	if !ok || resourceID == "" {
		return &plugins.ExecuteResponse{
			Result:  fmt.Sprintf("%s is required, but not provided", idKey),
			IsError: true,
		}, nil
	}

	path := fmt.Sprintf("/sites/%s/%s/%s", siteID, resource, resourceID)
	if err := p.DoGetRequest(ctx, path, dst); err != nil {
		return &plugins.ExecuteResponse{
			Result:  err.Error(),
			IsError: true,
		}, nil
	}

	return &plugins.ExecuteResponse{
		Result: dst,
	}, nil
}

// get performs an authenticated GET request to the UniFi Integration API and
// decodes the JSON response into dst.
func (p *UnifiToolPlugin) DoGetRequest(ctx context.Context, path string, dst any) error {
	url := LocalBaseUrl(p.driver.config.Address) + path
	p.driver.log.Debug("UniFi API request", "method", "GET", "path", path)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}
	SetHttpAPIHeaders(req, p.driver.config.APIKey)

	resp, err := p.client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("authentication failed (401): check api_key")
	}
	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("not found (404): %s", path)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d for %s", resp.StatusCode, path)
	}

	if err := json.NewDecoder(resp.Body).Decode(dst); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	return nil
}
