resource "azurerm_virtual_network" "kunlun_virtual_network" {
  name                = "${var.env_name}-vnet"
  resource_group_name = "${azurerm_resource_group.kunlun_resource_group.name}"
  address_space       = ["10.0.0.0/16"]
  location            = "${var.location}"
}
