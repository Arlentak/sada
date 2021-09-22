package docker_utils

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceInspect() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceInspectRead,
		Schema: map[string]*schema.Schema{
			"container_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"mounts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"destination": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"read_write": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"propagation": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gateway": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_prefixlen": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceInspectRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := i.(*ProviderConfig).DockerClient
	var container = data.Get("container_name").(string)

	res, err := client.ContainerInspect(ctx, container)
	if err != nil {
		return diag.FromErr(err)
	}

	var network []map[string]interface{}

	for key := range res.NetworkSettings.Networks {
		network = append(network, populateNetwork(res.NetworkSettings, key))
	}
	var mounts []map[string]interface{}

	for _, val := range res.Mounts {
		mounts = append(mounts, populateMount(val))
	}
	if err := data.Set("networks", network); err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("mounts", mounts); err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("environment", res.Config.Env); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(res.ID)

	return diags
}

func populateMount(mounts types.MountPoint) map[string]interface{} {
	return map[string]interface{}{
		"type":        mounts.Type,
		"source":      mounts.Source,
		"destination": mounts.Destination,
		"mode":        mounts.Mode,
		"read_write":  mounts.RW,
		"propagation": mounts.Propagation,
	}
}
func populateNetwork(data *types.NetworkSettings, network string) map[string]interface{} {
	return map[string]interface{}{
		"network_name": network,
		"ip_address":   data.Networks[network].IPAddress,
		"gateway":      data.Networks[network].Gateway,
		"ip_prefixlen": data.Networks[network].IPPrefixLen,
	}
}
