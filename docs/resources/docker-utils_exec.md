# docker-utils_exec

Provides a docker exec that run a command in a running container.

## Example Usage

```hcl
resource "docker-utils_exec" "exec" {
  container_name = "container_name"    
  commands = ["/bin/bash","-c","mkdir $FILENAME"] 
  environment = ["FILENAME=example"] 
}
```

## Argument Reference

The following arguments are supported:
    
  * `container_name` - (Required)  `string` The container name, alias or ID
  * `commands` - (Required) Specify the command that will be executed after creation.
  * `environment` - (Required) A list of env variables to Set in creation.
  * `destroy_commands` - (Optional) Specify the command that will be executed in destroy.
  * `destroy_environment` - (Optional) A list of env variables to Set in destroy.
  * `attach_stderr` - (Optional)  `boolean` attach STDERR (default false)
  * `attach_stdin` - (Optional)  `boolean` attach STDIN (default false)
  * `attach_stdout` - (Optional)  `boolean` attach STDOUT (default false)
  * `detach` - (Optional)  `boolean` Detached mode: run command in the background. (default false)
  * `tty` - (Optional)  `boolean` Allocate a pseudo-TTY (default false)
  * `privileged` - (Optional)  `boolean` Give extended privileges to the command (default false)
  * `working_dir` - (Optional)  `string` Working directory inside the container (default /)
  * `user` - (Optional) `string` Username or UID (format: <name|uid>[:<group|gid>]) (default root)

## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `key` - The API Key

## Import

Algolia API Key can be imported using the `key`

```shell
terraform import algolia_api_key.example {my algolia api key}
```