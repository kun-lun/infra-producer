resource "azurerm_resource_group" "kunlun_resource_group" {
  name     = "${var.env_name}-kunlun"
  location = "${var.location}"
}
