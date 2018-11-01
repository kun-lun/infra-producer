resource "azurerm_public_ip" "kunlun_devbox_public_ip" {
  name                         = "kunlun_devbox_public_ip"
  location                     = "${azurerm_resource_group.kunlun_resource_group.location}"
  resource_group_name          = "${azurerm_resource_group.kunlun_resource_group.name}"
  public_ip_address_allocation = "static"
}

resource "azurerm_network_interface" "kunlun_devbox_nic" {
  name                = "${var.env_name}-devbox-nic"
  location            = "${azurerm_resource_group.kunlun_resource_group.location}"
  resource_group_name = "${azurerm_resource_group.kunlun_resource_group.name}"

  ip_configuration {
    name                          = "${var.env_name}-devbox-nicip"
    subnet_id                     = "${azurerm_subnet.san_subnet.id}"
    private_ip_address_allocation = "dynamic"
    public_ip_address_id          = "${azurerm_public_ip.kunlun_devbox_public_ip.id}"
  }
}

resource "azurerm_virtual_machine" "kunlun_devbox" {
  name                  = "${var.env_name}-kunlun-devbox"
  location              = "${azurerm_resource_group.kunlun_resource_group.location}"
  resource_group_name   = "${azurerm_resource_group.kunlun_resource_group.name}"
  network_interface_ids = ["${azurerm_network_interface.kunlun_devbox_nic.id}"]
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
    name              = "kunlundevboxdisk"
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }

  os_profile {
    computer_name  = "devbox"
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
