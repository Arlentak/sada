package docker_utils

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDockerExec() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDockerExecCreate,
		ReadContext:   resourceDockerExecRead,
		UpdateContext: resourceDockerExecUpdate,
		DeleteContext: resourceDockerExecDelete,
		Importer: &schema.ResourceImporter{
			StateContext: importExecState,
		},
		Schema: map[string]*schema.Schema{
			"container_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"commands": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"destroy_commands": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"destroy_environment": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"environment": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attach_stderr": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"attach_stdin": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"attach_stdout": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"detach": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tty": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"privileged": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"working_dir": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDockerExecCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(*ProviderConfig).DockerClient

	var container = data.Get("container_name").(string)
	retContainer, err := client.ContainerInspect(ctx, container)
	if err != nil {
		return diag.Errorf("container :%s  not found", container)
	}

	var config types.ExecConfig

	var command = data.Get("commands").([]interface{})
	var environment = data.Get("environment").([]interface{})

	listOfCommands := make([]string, len(command), cap(command))

	listOfEnv := make([]string, len(environment), cap(environment))

	for i, v := range command {
		listOfCommands[i] = v.(string)
	}
	for i, v := range environment {
		listOfEnv[i] = v.(string)
	}

	config.Cmd = listOfCommands
	config.Env = listOfEnv

	config.Detach = data.Get("detach").(bool)
	config.Tty = data.Get("tty").(bool)
	config.Privileged = data.Get("privileged").(bool)
	config.WorkingDir = data.Get("working_dir").(string)
	config.User = data.Get("user").(string)
	config.AttachStderr = data.Get("attach_stderr").(bool)
	config.AttachStdin = data.Get("attach_stdin").(bool)
	config.AttachStdout = data.Get("attach_stdout").(bool)

	execId, err := client.ContainerExecCreate(ctx, retContainer.ID, config)
	if err != nil {
		return diag.Errorf("Could not create the docker exec err: %s", err)
	}
	data.SetId(execId.ID)
	var checkConfig types.ExecStartCheck
	checkConfig.Detach = true
	checkConfig.Tty = true
	err = client.ContainerExecStart(ctx, execId.ID, checkConfig)
	if err != nil {
		return diag.Errorf("Could not start the docker exec err")
	}
	return resourceDockerExecRead(ctx, data, i)
}
func resourceDockerExecUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(*ProviderConfig).DockerClient

	var container = data.Get("container_name").(string)

	retContainer, err := client.ContainerInspect(ctx, container)
	if err != nil {
		return diag.Errorf("container :%s  not found", container)
	}

	var config types.ExecConfig

	var command = data.Get("commands").([]interface{})
	var environment = data.Get("environment").([]interface{})

	listOfCommands := make([]string, len(command), cap(command))

	listOfEnv := make([]string, len(environment), cap(environment))

	for i, v := range command {
		listOfCommands[i] = v.(string)
	}
	for i, v := range environment {
		listOfEnv[i] = v.(string)
	}

	config.Cmd = listOfCommands
	config.Env = listOfEnv

	config.Detach = data.Get("detach").(bool)
	config.Tty = data.Get("tty").(bool)
	config.Privileged = data.Get("privileged").(bool)
	config.WorkingDir = data.Get("working_dir").(string)
	config.User = data.Get("user").(string)
	config.AttachStderr = data.Get("attach_stderr").(bool)
	config.AttachStdin = data.Get("attach_stdin").(bool)
	config.AttachStdout = data.Get("attach_stdout").(bool)

	execId, err := client.ContainerExecCreate(ctx, retContainer.ID, config)
	if err != nil {
		return diag.Errorf("Could not create the docker exec err: %s", err)
	}
	data.SetId(execId.ID)
	var checkConfig types.ExecStartCheck
	checkConfig.Detach = true
	checkConfig.Tty = true
	err = client.ContainerExecStart(ctx, execId.ID, checkConfig)
	if err != nil {
		return diag.Errorf("Could not start the docker exec err")
	}
	return resourceDockerExecRead(ctx, data, i)
}

func resourceDockerExecDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {

	client := i.(*ProviderConfig).DockerClient
	var container = data.Get("container_name").(string)
	retContainer, err := client.ContainerInspect(ctx, container)
	if err != nil {
		return diag.Errorf("container :%s  not found", container)
	}

	var config types.ExecConfig
	config.Detach = data.Get("detach").(bool)
	config.Tty = data.Get("tty").(bool)
	config.Privileged = data.Get("privileged").(bool)
	config.WorkingDir = data.Get("working_dir").(string)
	config.User = data.Get("user").(string)
	var command = data.Get("destroy_commands").([]interface{})
	var environment = data.Get("destroy_environment").([]interface{})
	listOfCommands := make([]string, len(command), cap(command))
	listOfEnvironment := make([]string, len(environment), cap(environment))
	for i, v := range command {
		listOfCommands[i] = v.(string)
	}
	for i, v := range environment {
		listOfEnvironment[i] = v.(string)
	}
	config.Cmd = listOfCommands
	config.Env = listOfEnvironment

	execId, err := client.ContainerExecCreate(ctx, retContainer.ID, config)
	if err != nil {
		return diag.Errorf("Could not destroy the docker exec err: %s", err)
	}
	var checkConfig types.ExecStartCheck
	checkConfig.Detach = false
	checkConfig.Tty = false
	err = client.ContainerExecStart(ctx, execId.ID, checkConfig)
	if err != nil {
		return diag.Errorf("Could not destroy the docker exec verify destroy commands")
	}
	return resourceDockerExecRead(ctx, data, i)

}

func resourceDockerExecRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics = nil
	return diags
}
func importExecState(ctx context.Context, data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	if err := data.Set("Id", data.Get("Id")); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{data}, nil
}
