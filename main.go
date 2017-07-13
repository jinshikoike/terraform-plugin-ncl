package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jinshikoike/terraform-ncl/ncl"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: ncl.Provider})
}
