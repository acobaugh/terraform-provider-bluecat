package main

import (
	"github.com/acobaugh/terraform-provider-bluecat/bluecat"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: bluecat.Provider,
	})
}
