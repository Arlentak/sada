package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/Kaginari/terraform-provider-docker-utils/docker-utils"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: docker_utils.Provider})
}