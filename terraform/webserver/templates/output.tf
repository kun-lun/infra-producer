output "environment" {
  value = "${var.azure_environment}"
}

output "resource_group_name" {
  value = "${azurerm_resource_group.kunlun_resource_group.name}"
}

output "vnet_name" {
  value = "${azurerm_virtual_network.kunlun_virtual_network.name}"
}
