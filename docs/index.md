# Docker-utils Provider

The Docker-utils Provider is used to interact with the many resources supported by Docker but not provided in docker provider. 
The provider needs to be configured with the proper host before it can be used.
Use the navigation to the left to read about the available resources.

## Example Usage
```hcl
terraform {
  required_version = ">= 0.13"

  required_providers {
    docker-utils = {
      source = "Kaginari/docker-utils"
    }
  }
}

variable "docker_tcp_host" {}
variable "docker_container_name" {}

provider "docker-utils" {
    host = var.docker_tcp_host
}
resource "docker-utils_exec" "exec" {

  container_name = var.docker_container_name    #(Required) the container alias or id

  # Exec Options

  attach_stderr = true    # optional default false
  attach_stdin = true     # optional default false
  attach_stdout = true    # optional default false
  detach  = true          # optional default false
  tty  = true             # optional default false
  privileged  = false     # optional default false
  user  = "root"          # optional default root
  working_dir = "/home"   # optional default root folder

  # Exec commands
  commands = ["/bin/bash","-c","mkdir example && touch example/$ME.txt"]                    # (Required) commands will be applied on apply
  destroy_commands = ["/bin/bash","-c","rm -rf example"]                                    # (Optional) commands will be applied on destroy

  # Exec environment
  environment = ["ME=example"]                                                              # (Required) environment will be applied on apply
  destroy_environment = []                                                                  # (Optional) environment will be applied on destroy
}

data "docker-utils_logs" "example" {
  container_name = "dev_mysql"
  tail = 1                                            #(Optional) Default all
  details = true                                      #(Optional) Default true
  show_stderr = true                                  #(Optional) Default true
  show_stdout = true                                  #(Optional) Default true
  timestamps = true                                   #(Optional) Default false
  from_date = "2020-10-18T06:32:35.587250Z"           #(Optional) Default null
  to_date = "2020-10-21T06:32:35.587250Z"             #(Optional) Default null

}

data "docker-utils_inspect" "example" {
  container_name =  var.docker_container_name
}
output "logs" {
  value = data.docker-utils_logs.example.logs
}

# List of container mounts 
output "container_mounts" {
  value = data.docker-utils_inspect.example.mounts
}
# List of container env variables 
output "container_environment" {
  value = data.docker-utils_inspect.example.environment
}
# List of container networks
output "container_networks" {
  value = data.docker-utils_inspect.example.networks
}
```

## Argument Reference

* `docker_tcp_host` - (Required) This is the address to the Docker host. If this is blank, the DOCKER_HOST environment variable will also be read.
* `docker_container_name` - (Required) This is the (ID , name or alias) of the docker container.