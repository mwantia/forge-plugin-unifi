package unifi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/mwantia/forge-plugin-unifi/internal/unifi/api"
	"github.com/mwantia/forge-sdk/pkg/plugins"
)

// apiResponse wraps the standard UniFi Integration API list envelope.
type apiResponse[T any] struct {
	Data       []T `json:"data"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	Count      int `json:"count"`
	TotalCount int `json:"totalCount"`
}

func (d *UnifiDriver) GetLifecycle() plugins.Lifecycle {
	return d
}

func (d *UnifiDriver) ListTools(_ context.Context, filter plugins.ListToolsFilter) (*plugins.ListToolsResponse, error) {
	if d.config == nil {
		return nil, fmt.Errorf("plugin not configured")
	}

	tools := make([]plugins.ToolDefinition, 0, len(d.config.Tools))
	for _, name := range d.config.Tools {
		def, ok := toolDefinitions[name]
		if !ok {
			d.log.Warn("Unknown tool in config, skipping", "tool", name)
			continue
		}
		if matchesFilter(def, filter) {
			tools = append(tools, def)
		}
	}

	return &plugins.ListToolsResponse{Tools: tools}, nil
}

func (d *UnifiDriver) GetTool(_ context.Context, name string) (*plugins.ToolDefinition, error) {
	if d.config == nil {
		return nil, fmt.Errorf("plugin not configured")
	}

	def, ok := toolDefinitions[name]
	if !ok {
		return nil, fmt.Errorf("tool %q not found", name)
	}

	if slices.Contains(d.config.Tools, name) {
		return &def, nil
	}
	return nil, fmt.Errorf("tool %q is not enabled", name)
}

func (d *UnifiDriver) Validate(_ context.Context, req plugins.ExecuteRequest) (*plugins.ValidateResponse, error) {
	if d.config == nil {
		return nil, fmt.Errorf("plugin not configured")
	}

	var errs []string
	requireStr := func(key string) {
		v, ok := req.Arguments[key]
		if !ok || v == "" {
			errs = append(errs, fmt.Sprintf("%q is required", key))
		}
	}

	switch req.Tool {
	case "list_local_sites":
		// no required parameters
	case "list_adopted_devices", "list_connected_clients", "list_networks",
		"list_wifi_broadcast", "list_firewall_zones", "list_firewall_policies":
		requireStr("site_id")
	case "get_adopted_device_details":
		requireStr("site_id")
		requireStr("device_id")
	case "get_connected_client_details":
		requireStr("site_id")
		requireStr("client_id")
	case "get_network_details":
		requireStr("site_id")
		requireStr("network_id")
	case "get_wifi_broadcast_details":
		requireStr("site_id")
		requireStr("wifi_broadcast_id")
	case "get_firewall_zone":
		requireStr("site_id")
		requireStr("firewall_zone_id")
	case "get_firewall_policy":
		requireStr("site_id")
		requireStr("firewall_policy_id")
	default:
		errs = append(errs, fmt.Sprintf("unknown tool %q", req.Tool))
	}

	return &plugins.ValidateResponse{Valid: len(errs) == 0, Errors: errs}, nil
}

func (d *UnifiDriver) Execute(ctx context.Context, req plugins.ExecuteRequest) (*plugins.ExecuteResponse, error) {
	if d.config == nil {
		return nil, fmt.Errorf("plugin not configured")
	}

	if !slices.Contains(d.config.Tools, req.Tool) {
		return &plugins.ExecuteResponse{
			Result:  fmt.Sprintf("tool '%s' is not enabled in unifi configuration", req.Tool),
			IsError: true,
		}, nil
	}

	args := req.Arguments
	switch req.Tool {
	case "list_local_sites":
		return d.execList(ctx, "/sites", args, func() any { return &apiResponse[api.LocalSite]{} },
			func(r any) any {
				res := r.(*apiResponse[api.LocalSite])
				return paginatedResult("sites", res.Data, res)
			})
	case "list_adopted_devices":
		return d.execListForSite(ctx, "devices", args, func() any { return &apiResponse[api.AdoptedDevice]{} },
			func(r any) any {
				res := r.(*apiResponse[api.AdoptedDevice])
				return paginatedResult("devices", res.Data, res)
			})
	case "get_adopted_device_details":
		return d.execGetForSite(ctx, "devices", "device_id", args, &api.AdoptedDeviceDetails{})
	case "list_connected_clients":
		return d.execListForSite(ctx, "clients", args, func() any { return &apiResponse[api.ConnectedClient]{} },
			func(r any) any {
				res := r.(*apiResponse[api.ConnectedClient])
				return paginatedResult("clients", res.Data, res)
			})
	case "get_connected_client_details":
		return d.execGetForSite(ctx, "clients", "client_id", args, &api.ConnectedClientDetails{})
	case "list_networks":
		return d.execListForSite(ctx, "networks", args, func() any { return &apiResponse[api.Network]{} },
			func(r any) any {
				res := r.(*apiResponse[api.Network])
				return paginatedResult("networks", res.Data, res)
			})
	case "get_network_details":
		return d.execGetForSite(ctx, "networks", "network_id", args, &api.NetworkDetails{})
	case "list_wifi_broadcast":
		return d.execListForSite(ctx, "wifi/broadcasts", args, func() any { return &apiResponse[api.WifiBroadcast]{} },
			func(r any) any {
				res := r.(*apiResponse[api.WifiBroadcast])
				return paginatedResult("wifi_broadcasts", res.Data, res)
			})
	case "get_wifi_broadcast_details":
		return d.execGetForSite(ctx, "wifi/broadcasts", "wifi_broadcast_id", args, &api.WifiBroadcastDetails{})
	case "list_firewall_zones":
		return d.execListForSite(ctx, "firewall/zones", args, func() any { return &apiResponse[api.FirewallZone]{} },
			func(r any) any {
				res := r.(*apiResponse[api.FirewallZone])
				return paginatedResult("firewall_zones", res.Data, res)
			})
	case "get_firewall_zone":
		return d.execGetForSite(ctx, "firewall/zones", "firewall_zone_id", args, &api.FirewallZone{})
	case "list_firewall_policies":
		return d.execListForSite(ctx, "firewall/policies", args, func() any { return &apiResponse[api.FirewallPolicy]{} },
			func(r any) any {
				res := r.(*apiResponse[api.FirewallPolicy])
				return paginatedResult("firewall_policies", res.Data, res)
			})
	case "get_firewall_policy":
		return d.execGetForSite(ctx, "firewall/policies", "firewall_policy_id", args, &api.FirewallPolicy{})
	default:
		return &plugins.ExecuteResponse{
			Result:  fmt.Sprintf("unknown tool: %s", req.Tool),
			IsError: true,
		}, nil
	}
}

// execList fetches a paginated list from path, decodes into newResult(), and shapes the output via shape().
func (d *UnifiDriver) execList(ctx context.Context, path string, args map[string]any, newResult func() any, shape func(any) any) (*plugins.ExecuteResponse, error) {
	dst := newResult()
	if err := d.get(ctx, path+paginationQuery(args), dst); err != nil {
		return &plugins.ExecuteResponse{Result: err.Error(), IsError: true}, nil
	}
	return &plugins.ExecuteResponse{Result: shape(dst)}, nil
}

// execListForSite fetches a paginated list scoped to a site.
func (d *UnifiDriver) execListForSite(ctx context.Context, resource string, args map[string]any, newResult func() any, shape func(any) any) (*plugins.ExecuteResponse, error) {
	siteID, ok := args["site_id"].(string)
	if !ok || siteID == "" {
		return &plugins.ExecuteResponse{Result: "site_id is required", IsError: true}, nil
	}
	return d.execList(ctx, fmt.Sprintf("/sites/%s/%s", siteID, resource), args, newResult, shape)
}

// execGetForSite fetches a single resource by ID, scoped to a site.
// dst must be a pointer to the target struct.
func (d *UnifiDriver) execGetForSite(ctx context.Context, resource, idKey string, args map[string]any, dst any) (*plugins.ExecuteResponse, error) {
	siteID, ok := args["site_id"].(string)
	if !ok || siteID == "" {
		return &plugins.ExecuteResponse{Result: "site_id is required", IsError: true}, nil
	}
	resourceID, ok := args[idKey].(string)
	if !ok || resourceID == "" {
		return &plugins.ExecuteResponse{Result: fmt.Sprintf("%s is required", idKey), IsError: true}, nil
	}

	path := fmt.Sprintf("/sites/%s/%s/%s", siteID, resource, resourceID)
	if err := d.get(ctx, path, dst); err != nil {
		return &plugins.ExecuteResponse{Result: err.Error(), IsError: true}, nil
	}
	return &plugins.ExecuteResponse{Result: dst}, nil
}

// paginatedResult builds the standard paginated response map.
func paginatedResult[T any](key string, data []T, r *apiResponse[T]) map[string]any {
	return map[string]any{
		key:           data,
		"count":       r.Count,
		"total_count": r.TotalCount,
		"offset":      r.Offset,
		"limit":       r.Limit,
	}
}

// paginationQuery builds a query string from optional limit/offset args.
func paginationQuery(args map[string]any) string {
	limit, hasLimit := intArg(args, "limit")
	offset, hasOffset := intArg(args, "offset")
	if !hasLimit && !hasOffset {
		return ""
	}
	q := "?"
	if hasLimit {
		q += fmt.Sprintf("limit=%d", limit)
	}
	if hasOffset {
		if hasLimit {
			q += "&"
		}
		q += fmt.Sprintf("offset=%d", offset)
	}
	return q
}

func intArg(args map[string]any, key string) (int, bool) {
	v, ok := args[key]
	if !ok || v == nil {
		return 0, false
	}
	switch n := v.(type) {
	case float64:
		return int(n), true
	case int:
		return n, true
	}
	return 0, false
}

// get performs an authenticated GET request to the UniFi Integration API and
// decodes the JSON response into dst.
func (d *UnifiDriver) get(ctx context.Context, path string, dst any) error {
	url := d.baseURL() + path
	d.log.Debug("UniFi API request", "method", "GET", "path", path)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}
	d.setHeaders(req)

	resp, err := d.client.Do(req)
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

func matchesFilter(def plugins.ToolDefinition, f plugins.ListToolsFilter) bool {
	if def.Deprecated && !f.Deprecated {
		return false
	}
	if f.Prefix != "" && !strings.HasPrefix(def.Name, f.Prefix) {
		return false
	}
	if len(f.Tags) > 0 {
		for _, want := range f.Tags {
			for _, have := range def.Tags {
				if have == want {
					goto tagMatched
				}
			}
		}
		return false
	tagMatched:
	}
	return true
}
