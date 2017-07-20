package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jinshikoike/terraform-provider-ncl/ncl"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: ncl.Provider})
}
