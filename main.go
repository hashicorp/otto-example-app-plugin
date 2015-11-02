package main

import (
	"github.com/hashicorp/otto/plugin"
)

func main() {
	// This runs this binary as an Otto plugin. No other machinery is needed
	// to make the plugin work!
	plugin.Serve(&plugin.ServeOpts{
		AppFunc: AppFactory,
	})
}
