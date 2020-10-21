package docker_utils

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DOCKER_HOST", "unix:///var/run/docker.sock"),
				Description: "The Docker daemon address",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"docker-utils_exec": resourceDockerExec(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"docker-utils_logs":    dataSourceLogs(),
			"docker-utils_inspect": dataSourceInspect(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		Host: d.Get("host").(string),
	}

	client, err := config.NewClient()
	if err != nil {
		diag.Errorf("Error initializing Docker client: %s", err)
	}

	_, err = client.Ping(ctx)
	if err != nil {
		diag.Errorf("Error pinging Docker server: %s", err)
	}

	providerConfig := ProviderConfig{
		DockerClient: client,
	}

	return &providerConfig, nil
}
