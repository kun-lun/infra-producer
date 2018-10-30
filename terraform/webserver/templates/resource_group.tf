resource "azurerm_resource_group" "jindou_resource_group" {
  name     = "${var.resource_group_prefix}-${var.env_name}-jindou"
  location = "${var.location}"
}
