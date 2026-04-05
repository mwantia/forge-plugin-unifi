package unifi

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/mitchellh/mapstructure"
	"github.com/mwantia/forge-sdk/pkg/plugins"
)

const (
	PluginName        = "unifi"
	PluginAuthor      = "forge"
	PluginVersion     = "0.1.0"
	PluginDescription = "UniFi Network integration for querying local network infrastructure via the UniFi Integration API"
)

type UnifiDriver struct {
	plugins.UnimplementedDriver
	log    hclog.Logger
	config *UnifiConfig
	client *http.Client
}

type UnifiConfig struct {
	Address         string `mapstructure:"address"`
	APIKey          string `mapstructure:"api_key"`
	Timeout         string `mapstructure:"timeout"`
	InsecureSkipTLS bool   `mapstructure:"insecure_skip_tls"`
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

	url := LocalBaseUrl(d.config.Address) + "/sites"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false, fmt.Errorf("failed to build probe request: %w", err)
	}
	SetHttpAPIHeaders(req, d.config.APIKey)

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

	d.config = cfg

	d.log.Info("UniFi configured", "address", cfg.Address, "insecure_skip_tls", cfg.InsecureSkipTLS)
	return nil
}

func (d *UnifiDriver) GetToolsPlugin(ctx context.Context) (plugins.ToolsPlugin, error) {
	return NewUnifiToolPlugin(d)
}
