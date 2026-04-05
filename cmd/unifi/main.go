package main

import (
	"github.com/mwantia/forge-plugin-unifi/internal/unifi"
	"github.com/mwantia/forge-sdk/pkg/plugins/grpc"
)

func main() {
	grpc.Serve(unifi.NewUnifiDriver)
}
