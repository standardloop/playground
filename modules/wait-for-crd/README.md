<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
| ---- | ------- |
| <a name="provider_null"></a> [null](#provider\_null) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
| ---- | ---- |
| [null_resource.wait_for_crds](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) | resource |

## Inputs

| Name | Description | Type | Default | Required |
| ---- | ----------- | ---- | ------- | :------: |
| <a name="input_attempts"></a> [attempts](#input\_attempts) | n/a | `number` | `30` | no |
| <a name="input_crds"></a> [crds](#input\_crds) | n/a | `list(string)` | `[]` | no |
| <a name="input_timeout"></a> [timeout](#input\_timeout) | n/a | `string` | `"180s"` | no |

## Outputs

No outputs.
<!-- END_TF_DOCS -->
