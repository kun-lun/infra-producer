resource "azurerm_subnet" "san_subnet" {
  name                 = "${var.env_name}-subnet"
  resource_group_name  = "${azurerm_resource_group.kunlun_resource_group.name}"
  virtual_network_name = "${azurerm_virtual_network.kunlun_virtual_network.name}"
  address_prefix       = "${cidrsubnet(azurerm_virtual_network.kunlun_virtual_network.address_space[0], 8, 0)}"
}