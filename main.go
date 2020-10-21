package main

import (

	"github.com/Kaginari/terraform-provider-docker-utils/docker-utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: docker_utils.Provider})
}
