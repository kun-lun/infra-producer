
resource "azurerm_network_interface" "kunlun_nic" {
  name                = "${var.env_name}-nic-${count.index}"
  count               = "${var.web_server_vm_count}"
  location            = "${azurerm_resource_group.kunlun_resource_group.location}"
  resource_group_name = "${azurerm_resource_group.kunlun_resource_group.name}"

  ip_configuration {
    name                          = "${var.env_name}-nicip-${count.index}"
    subnet_id                     = "${azurerm_subnet.san_subnet.id}"
    private_ip_address_allocation = "dynamic"
  }
}

resource "azurerm_virtual_machine" "kunlun_vm" {
  name                  = "${var.env_name}-vm-${count.index}"
  location              = "${azurerm_resource_group.kunlun_resource_group.location}"
  resource_group_name   = "${azurerm_resource_group.kunlun_resource_group.name}"
  network_interface_ids = ["${azurerm_network_interface.kunlun_nic.*.id[count.index]}"]
  vm_size               = "${var.web_server_vm_size}"
  count                 = "${var.web_server_vm_count}"


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
    name              = "kunlunvmdisk${count.index}"
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }

  os_profile {
    computer_name  = "hostname${count.index}"
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

resource "azurerm_network_interface_backend_address_pool_association" "kunlun_vm_nic_backend_pool_association" {
  count                   = "${var.web_server_vm_count}"
  network_interface_id    = "${azurerm_network_interface.kunlun_nic.*.id[count.index]}"
  ip_configuration_name   = "${var.env_name}-nicip-${count.index}"
  backend_address_pool_id = "${azurerm_lb_backend_address_pool.kunlun_load_balancer_backend_address_pool.id}"
}
