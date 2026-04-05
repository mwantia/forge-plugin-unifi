package unifi

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/mitchellh/mapstructure"
	"github.com/mwantia/forge-sdk/pkg/errors"
	"github.com/mwantia/forge-sdk/pkg/plugins"
)

const (
	PluginName        = "unifi"
	PluginAuthor      = "forge"
	PluginVersion     = "0.1.0"
	PluginDescription = "UniFi Network integration for querying local network infrastructure via the UniFi Integration API"
)

type UnifiDriver struct {
	plugins.UnimplementedToolsPlugin
	log    hclog.Logger
	config *UnifiConfig
	client *http.Client
}

type UnifiConfig struct {
	Address         string   `mapstructure:"address"`
	APIKey          string   `mapstructure:"api_key"`
	Timeout         string   `mapstructure:"timeout"`
	InsecureSkipTLS bool     `mapstructure:"insecure_skip_tls"`
	Tools           []string `mapstructure:"tools"`
}

func NewUnifiDriver(log hclog.Logger) plugins.Driver {
	return &UnifiDriver{
		log: log.Named(PluginName),
	}
}

func (d *UnifiDriver) GetPluginInfo() plugins.PluginInfo {
	return plugins.PluginInfo{
		Name:        PluginName,
		Author:      PluginAuthor,
		Version:     PluginVersion,
		Description: PluginDescription,
	}
}

func (d *UnifiDriver) ProbePlugin(ctx context.Context) (bool, error) {
	if d.config == nil {
		return false, fmt.Errorf("plugin not configured")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, d.baseURL()+"/sites", nil)
	if err != nil {
		return false, fmt.Errorf("failed to build probe request: %w", err)
	}
	d.setHeaders(req)

	resp, err := d.client.Do(req)
	if err != nil {
		return false, fmt.Errorf("probe request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return false, fmt.Errorf("authentication failed: check api_key")
	}
	return resp.StatusCode == http.StatusOK, nil
}

func (d *UnifiDriver) GetCapabilities(ctx context.Context) (*plugins.DriverCapabilities, error) {
	return &plugins.DriverCapabilities{
		Types: []string{plugins.PluginTypeTools},
		Tools: &plugins.ToolsCapabilities{
			SupportsAsyncExecution: false,
		},
	}, nil
}

func (d *UnifiDriver) OpenDriver(ctx context.Context) error {
	return nil
}

func (d *UnifiDriver) CloseDriver(ctx context.Context) error {
	return nil
}

func (d *UnifiDriver) ConfigDriver(ctx context.Context, config plugins.PluginConfig) error {
	cfg := &UnifiConfig{}
	if err := mapstructure.Decode(config.ConfigMap, cfg); err != nil {
		return fmt.Errorf("failed to decode config: %w", err)
	}

	if cfg.Address == "" {
		return fmt.Errorf("address is required")
	}
	cfg.Address = strings.TrimRight(cfg.Address, "/")
	if cfg.APIKey == "" {
		return fmt.Errorf("api_key is required")
	}
	if cfg.Timeout == "" {
		cfg.Timeout = "30s"
	}
	timeout, err := time.ParseDuration(cfg.Timeout)
	if err != nil {
		return fmt.Errorf("invalid timeout %q: %w", cfg.Timeout, err)
	}
	if len(cfg.Tools) == 0 {
		cfg.Tools = defaultTools()
	}

	d.config = cfg
	d.client = &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: cfg.InsecureSkipTLS, //nolint:gosec
			},
		},
	}

	d.log.Info("UniFi configured", "address", cfg.Address, "tools", cfg.Tools, "insecure_skip_tls", cfg.InsecureSkipTLS)
	return nil
}

func (d *UnifiDriver) GetProviderPlugin(ctx context.Context) (plugins.ProviderPlugin, error) {
	return nil, errors.ErrPluginNotSupported
}

func (d *UnifiDriver) GetMemoryPlugin(ctx context.Context) (plugins.MemoryPlugin, error) {
	return nil, errors.ErrPluginNotSupported
}

func (d *UnifiDriver) GetChannelPlugin(ctx context.Context) (plugins.ChannelPlugin, error) {
	return nil, errors.ErrPluginNotSupported
}

func (d *UnifiDriver) GetToolsPlugin(ctx context.Context) (plugins.ToolsPlugin, error) {
	return d, nil
}

func (d *UnifiDriver) GetSandboxPlugin(_ context.Context) (plugins.SandboxPlugin, error) {
	return nil, errors.ErrPluginNotSupported
}

func (d *UnifiDriver) baseURL() string {
	return d.config.Address + "/proxy/network/integration/v1"
}

func (d *UnifiDriver) setHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-API-Key", d.config.APIKey)
}

func defaultTools() []string {
	tools := make([]string, 0, len(toolDefinitions))
	for name := range toolDefinitions {
		tools = append(tools, name)
	}
	return tools
}
