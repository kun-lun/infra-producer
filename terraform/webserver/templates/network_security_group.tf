resource "azurerm_network_security_group" "kunlun_jumpbox_network_security_group" {
  name                = "${var.env_name}-jumpbox-nsg"
  location            = "${azurerm_resource_group.kunlun_resource_group.location}"
  resource_group_name = "${azurerm_resource_group.kunlun_resource_group.name}"
}

resource "azurerm_network_security_rule" "kunlun_jumpbox_network_security_rule_ssh" {
  name                        = "Allow-Ssh"
  priority                    = 100
  direction                   = "Inbound"
  access                      = "Allow"
  protocol                    = "Tcp"
  source_port_range           = "*"
  destination_port_range      = "22"
  source_address_prefix       = "*"
  destination_address_prefix  = "*"
  resource_group_name         = "${azurerm_resource_group.kunlun_resource_group.name}"
  network_security_group_name = "${azurerm_network_security_group.kunlun_jumpbox_network_security_group.name}"
}

resource "azurerm_network_security_group" "kunlun_server_network_security_group" {
  name                = "${var.env_name}-server-nsg"
  location            = "${azurerm_resource_group.kunlun_resource_group.location}"
  resource_group_name = "${azurerm_resource_group.kunlun_resource_group.name}"
}

resource "azurerm_network_security_rule" "kunlun_server_network_security_rule_http" {
  name                        = "Allow-Http"
  priority                    = 100
  direction                   = "Inbound"
  access                      = "Allow"
  protocol                    = "Tcp"
  source_port_range           = "*"
  destination_port_range      = "80"
  source_address_prefix       = "*"
  destination_address_prefix  = "*"
  resource_group_name         = "${azurerm_resource_group.kunlun_resource_group.name}"
  network_security_group_name = "${azurerm_network_security_group.kunlun_server_network_security_group.name}"
}