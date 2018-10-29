output "environment" {
  value = "${var.azure_environment}"
}

output "resource_group_name" {
  value = "${azurerm_resource_group.jindou_resource_group.name}"
}

output "vnet_name" {
  value = "${azurerm_virtual_network.jindou_virtual_network.name}"
}

output "subnet_name" {
  value = "${azurerm_subnet.jumpbox.name}"
}

output "internal_cidr" {
  value = "${azurerm_subnet.jumpbox.address_prefix}"
}

output "internal_gw" {
  value = "${cidrhost(azurerm_subnet.jumpbox.address_prefix, 1)}"
}

output "internal_ip" {
  value = "${cidrhost(azurerm_subnet.jumpbox.address_prefix, 6)}"
}
