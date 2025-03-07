---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cloudconnexa_network_ip_service Resource - terraform-provider-cloudconnexa"
subcategory: ""
description: |-
  
---

# cloudconnexa_network_ip_service (Resource)



## Example Usage

```terraform
resource "cloudconnexa_network" "this" {
  description     = "Test network"
  egress          = true
  name            = "my_test_network"
  internet_access = "SPLIT_TUNNEL_ON"
}

resource "cloudconnexa_route" "this" {
  description     = "Managed by Terraform"
  type            = "IP_V4"
  network_item_id = cloudconnexa_network.this.id
  subnet          = "192.168.144.0/24"
}

resource "cloudconnexa_network_ip_service" "example1" {
  name        = "example1"
  description = "Managed by Terraform"
  type        = "IP_SOURCE"
  network_id  = cloudconnexa_network.this.id
  routes      = ["192.168.144.1/32"]
  config {
    service_types = ["ANY"]
  }

  depends_on = [cloudconnexa_route.this]
}

resource "cloudconnexa_network_ip_service" "example2" {
  name        = "example2"
  description = "Managed by Terraform"
  type        = "IP_SOURCE"
  network_id  = cloudconnexa_network.this.id
  routes      = ["192.168.144.2/32"]
  config {
    service_types = ["HTTP", "HTTPS"]
    custom_service_types {
      protocol = "TCP" # all TCP ports
    }
    custom_service_types {
      protocol  = "UDP"
      from_port = 1194
      to_port   = 1194
    }
    custom_service_types {
      protocol  = "UDP"
      from_port = 5000
      to_port   = 5010
    }
    custom_service_types {
      protocol  = "ICMP"
      from_port = 8
      to_port   = 8
    }
    custom_service_types {
      protocol  = "ICMP"
      from_port = 20
      to_port   = 22
    }
  }

  depends_on = [cloudconnexa_route.this]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `network_id` (String)
- `type` (String)

### Optional

- `config` (Block List, Max: 1) (see [below for nested schema](#nestedblock--config))
- `description` (String)
- `routes` (List of String)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--config"></a>
### Nested Schema for `config`

Optional:

- `custom_service_types` (Block Set) (see [below for nested schema](#nestedblock--config--custom_service_types))
- `service_types` (List of String)

<a id="nestedblock--config--custom_service_types"></a>
### Nested Schema for `config.custom_service_types`

Required:

- `protocol` (String)

Optional:

- `from_port` (Number)
- `to_port` (Number)

## Import

Import is supported using the following syntax:

```shell
terraform import cloudconnexa_network_ip_service.example <id>
```
