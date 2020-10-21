# docker-utils_inspect

Use this data source to access information about an existing docker container.

## Example Usage

```hcl
data "docker-utils_inspect" "example" {
  container_name = "container_name"
}
```

## Argument Reference

The following arguments are supported:
* `name` - (Required) The name, alias or ID of the container.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `id` - The id of the container.
* `environment` - A List of string represents the env variables in the container.
* `mounts` - A List of mounts block as documented below.
* `networks` - A List of networks block as documented below.

---

A `mounts` block exports the following:
* `source` - The source of the mount. For bind mounts, this is the path to the file or directory on the Docker daemon host.
* `destination` - the destination takes as its value the path where the file or directory is mounted in the container.
* `propagation` - The bind-propagation option, if present, changes the bind propagation. May be one of rprivate, private, rshared, shared, rslave, slave.
* `read_write` - The read_write option, boolean, if `false` the bind mount is mounted into the container as read-only, else if `true`  the bind mount is mounted into the container as read-write
* `type` - The type of the mount, which can be bind, volume, or tmpfs.
* `mode` - The mode is a string that represents how the volume is mounted expl: 'RW' - read write .


A `networks` block exports the following:
* `network_name` - The name of the network 
* `ip_address` - The container's ip address in this network.
* `gateway` - The gateway address of the network.
