package main

import (
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/pkg/plugins/ipblacklist"
	"go.minekube.com/gate/pkg/plugins/configdownloader"
)

func main() {
	// Register plugins
	proxy.Plugins = append(proxy.Plugins,
		ipblacklist.Plugin,
		configdownloader.Plugin,
	)
	
	// Start Gate
	gate.Execute()
}
