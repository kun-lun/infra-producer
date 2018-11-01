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

resource "azurerm_public_ip" "kunlun_jumpbox_public_ip" {
  name                         = "${var.env_name}-jumpbox-public-ip"
  location                     = "${azurerm_resource_group.kunlun_resource_group.location}"
  resource_group_name          = "${azurerm_resource_group.kunlun_resource_group.name}"
  public_ip_address_allocation = "static"
}

resource "azurerm_network_interface" "kunlun_jumpbox_nic" {
  name                      = "${var.env_name}-jumpbox-nic"
  location                  = "${azurerm_resource_group.kunlun_resource_group.location}"
  resource_group_name       = "${azurerm_resource_group.kunlun_resource_group.name}"
  network_security_group_id = "${azurerm_network_security_group.kunlun_jumpbox_network_security_group.id}"
  ip_configuration {
    name                          = "${var.env_name}-jumpbox-nicip"
    subnet_id                     = "${azurerm_subnet.san_subnet.id}"
    private_ip_address_allocation = "dynamic"
    public_ip_address_id          = "${azurerm_public_ip.kunlun_jumpbox_public_ip.id}"
  }
}

resource "azurerm_virtual_machine" "kunlun_jumpbox" {
  name                  = "${var.env_name}-jumpbox"
  location              = "${azurerm_resource_group.kunlun_resource_group.location}"
  resource_group_name   = "${azurerm_resource_group.kunlun_resource_group.name}"
  network_interface_ids = ["${azurerm_network_interface.kunlun_jumpbox_nic.id}"]
  vm_size               = "Standard_DS2_v2"


  # Comment this line to keep the OS disk after deleting the VM
  delete_os_disk_on_termination = true

  # Comment this line to keep the data disks after deleting the VM
  delete_data_disks_on_termination = true

  storage_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }

  storage_os_disk {
    name              = "${var.env_name}-jumpbox-disk"
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }

  os_profile {
    computer_name  = "jumpbox"
    admin_username = "testadmin"
    admin_password = "Password1234!"
  }

  os_profile_linux_config {
    disable_password_authentication = false
  }

  tags {
    environment = "staging"
  }
}
