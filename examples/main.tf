terraform {
  required_version = ">= 0.13"

  required_providers {
    docker-utils = {
      source = "Kaginari/docker-utils"
      version = "9.9.9"
    }

  }
}

variable "docker_container_name" {
  default = "proxy"
}


provider "docker-utils" {
  host = "tcp://127.0.0.1:2376"
}
resource "docker-utils_exec" "create_ssl" {
  container_name = "proxy"
  attach_stderr = false
  attach_stdin = false
  attach_stdout = false
  detach  = true
  tty  = true
  commands = ["/bin/bash","-c","ls"]
}

//resource "docker-utils_exec" "exec" {
//
//  container_name = var.docker_container_name    #(Required) the container alias or id
//
//  # Exec Options
//
//  attach_stderr = true    # optional default false
//  attach_stdin = true     # optional default false
//  attach_stdout = true    # optional default false
//  detach  = true          # optional default false
//  tty  = true             # optional default false
//  privileged  = false     # optional default false
//  user  = "root"          # optional default root
//  working_dir = "/home"   # optional default root folder
//
//  # Exec commands
//  commands = ["/bin/bash","-c","mkdir example && touch example/$ME.txt"]                    # (Required) commands will be applied on apply
//  destroy_commands = ["/bin/bash","-c","rm -rf example"]                                    # (Optional) commands will be applied on destroy
//
//  # Exec environment
//  environment = ["ME=example"]                                                              # (Required) environment will be applied on apply
//  destroy_environment = []                                                                  # (Optional) environment will be applied on destroy
//}

//data "docker-utils_logs" "example" {
//  container_name = "proxy"
//  tail = 1                                            #(Optional) Default all
//  details = true                                      #(Optional) Default true
//  show_stderr = true                                  #(Optional) Default true
//  show_stdout = true                                  #(Optional) Default true
//  timestamps = true                                   #(Optional) Default false
//  from_date = "2020-10-18T06:32:35.587250Z"           #(Optional) Default null
//  to_date = "2020-10-21T06:32:35.587250Z"             #(Optional) Default null
//
//}

//data "docker-utils_inspect" "example" {
//  container_name =  var.docker_container_name
//}
//output "logs" {
//  value = data.docker-utils_logs.example.logs
//}
//
//# List of container mounts
//output "container_mounts" {
//  value = data.docker-utils_inspect.example.mounts
//}
//# List of container env variables
//output "container_environment" {
//  value = data.docker-utils_inspect.example.environment
//}
//# List of container networks
//output "container_networks" {
//  value = data.docker-utils_inspect.example.networks
//}
//output "container_id" {
//  value = data.docker-utils_inspect.example.id
//}
