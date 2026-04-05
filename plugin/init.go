package plugin

import (
	"github.com/mwantia/forge-plugin-unifi/internal/unifi"
	"github.com/mwantia/forge-sdk/pkg/plugins"
)

func init() {
	plugins.Register(unifi.PluginName, unifi.PluginDescription, unifi.NewUnifiDriver)
}
