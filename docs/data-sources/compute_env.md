---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nftower_compute_env Data Source - terraform-provider-nftower"
subcategory: ""
description: |-
  Example data source
---

# nftower_compute_env (Data Source)

Example data source



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) The ID of this resource.

### Optional

- **config** (Attributes) (see [below for nested schema](#nestedatt--config))
- **name** (String)

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

- **CliPath** (String)
- **ComputeJobRole** (String)
- **Discriminator** (String)
- **Forge** (String)
- **HeadJobCpus** (String)
- **HeadJobMemoryMb** (String)
- **HeadJobRole** (String)
- **HeadQueue** (String)
- **PostRunScript** (String)
- **PreRunScript** (String)
- **WorkDir** (String)
- **compute_queue** (String)
- **region** (String)

