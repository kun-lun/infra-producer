resource "azurerm_resource_group" "jindou_resource_group" {
  name     = "${var.env_name}-jindou"
  location = "${var.location}"
}
