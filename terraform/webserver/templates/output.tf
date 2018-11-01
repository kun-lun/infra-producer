output "environment" {
  value = "${var.azure_environment}"
}

output "resource_group_name" {
  value = "${azurerm_resource_group.jindou_resource_group.name}"
}

output "vnet_name" {
  value = "${azurerm_virtual_network.jindou_virtual_network.name}"
}
