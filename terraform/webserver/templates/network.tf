# Create a virtual network in the default resource group
resource "azurerm_virtual_network" "jindou_virtual_network" {
  name                = "jindou"
  resource_group_name = "${azurerm_resource_group.jindou_resource_group.name}"
  address_space       = ["10.0.0.0/16"]
  location            = "${var.location}"
}

resource "azurerm_subnet" "jumpbox" {
  name                 = "jumpbox"
  resource_group_name  = "${azurerm_resource_group.jindou_resource_group.name}"
  virtual_network_name = "${azurerm_virtual_network.jindou_virtual_network.name}"
  address_prefix       = "${cidrsubnet(azurerm_virtual_network.jindou_virtual_network.address_space[0], 8, 0)}"
}

resource "azurerm_subnet" "jindou_subnets" {
  name                 = "subnet-${count.index + 1}"
  resource_group_name  = "${azurerm_resource_group.jindou_resource_group.name}"
  virtual_network_name = "${azurerm_virtual_network.jindou_virtual_network.name}"
  address_prefix       = "${cidrsubnet(azurerm_virtual_network.jindou_virtual_network.address_space[0], 8, count.index + 1)}"

  count = 2
}
