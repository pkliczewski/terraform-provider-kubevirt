package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/pkliczewski/terraform-provider-kubevirt/kubevirt"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kubevirt.Provider})
}
