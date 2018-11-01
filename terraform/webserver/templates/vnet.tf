resource "azurerm_virtual_network" "jindou_virtual_network" {
  name                = "jindou"
  resource_group_name = "${azurerm_resource_group.jindou_resource_group.name}"
  address_space       = ["10.0.0.0/16"]
  location            = "${var.location}"
}
