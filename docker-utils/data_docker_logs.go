package docker_utils

import (
	"context"
	"io/ioutil"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/docker/docker/api/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLogs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLogsRead,
		Schema: map[string]*schema.Schema{
			"container_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"details": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"show_stderr": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"show_stdout": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tail": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "all",
			},
			"timestamps": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"from_date": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"to_date": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"logs": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceLogsRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := i.(*ProviderConfig).DockerClient

	var container = data.Get("container_name").(string)
	var logsOptions types.ContainerLogsOptions
	retContainer, err := client.ContainerInspect(ctx, container)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(data.Get("from_date").(string)) != 0 {
		sinceTime, err := time.Parse(time.RFC3339, data.Get("from_date").(string))
		if err != nil {
			return diag.Errorf("check from_date format must be RFC3339 error : %s", err)
		}
		logsOptions.Since = sinceTime.Format(time.RFC3339)
	}
	if len(data.Get("to_date").(string)) != 0 {
		untilTime, err := time.Parse(time.RFC3339, data.Get("to_date").(string))
		if err != nil {
			return diag.Errorf("check to_date format must be RFC3339 error : %s", err)
		}
		logsOptions.Since = untilTime.Format(time.RFC3339)
	}

	logsOptions.Details = data.Get("details").(bool)

	// Follow must be false otherwise it link logs to run

	logsOptions.Follow = false
	logsOptions.ShowStderr = data.Get("show_stderr").(bool)
	logsOptions.ShowStdout = data.Get("show_stdout").(bool)
	logsOptions.Tail = data.Get("tail").(string)
	logsOptions.Timestamps = data.Get("timestamps").(bool)

	logs, err := client.ContainerLogs(ctx, retContainer.ID, logsOptions)
	if err != nil {
		return diag.Errorf("can't get logs %s", err)
	}
	var bodyBytes []byte
	if logs != nil {
		bodyBytes, _ = ioutil.ReadAll(logs)
	}
	bodyString := string(bodyBytes)
	if !utf8.ValidString(bodyString) {
		v := make([]rune, 0, len(bodyString))
		for i, r := range bodyString {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(bodyString[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		bodyString = string(v)
	}

	if err := data.Set("logs", bodyString); err != nil {
		return diag.FromErr(err)
	}
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
