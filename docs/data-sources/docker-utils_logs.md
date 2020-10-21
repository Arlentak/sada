# docker-utils_logs

Use this data source to access logs of an existing docker container.

## Example Usage

```hcl
data "docker-utils_logs" "example" {
  container_name = "container_name"
}
```

## Argument Reference

The following arguments are supported:
* `name` - (Required) The name, alias or ID of the container.
* `details` - (Optional) Show extra details provided to logs. (default true) 
* `show_stderr` - (Optional)  `boolean` is typically used to output error messages.(default true)
* `show_stdout` - (Optional) `boolean` is typically used to output screen messages.(default true)
* `tail` - (Optional) `int` Number of lines to show from the end of the logs.(default all)
* `timestamps` - (Optional) `boolean` Show timestamps.(default false)
* `from_date` - (Optional)  `string` Show logs since timestamp (e.g. 2020-10-18T13:23:37).
* `to_date` - (Optional) `string` Show logs before a timestamp (e.g. 2020-10-18T16:23:37).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `logs` - The id of the container.


---

