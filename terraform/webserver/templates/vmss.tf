resource "azurerm_lb_nat_pool" "kunlun_vmss_nat_pool" {
  resource_group_name            = "${azurerm_resource_group.kunlun_resource_group.name}"
  name                           = "kunlun-ssh"
  loadbalancer_id                = "${azurerm_lb.kunlun_load_balancer.id}"
  protocol                       = "Tcp"
  frontend_port_start            = 50000
  frontend_port_end              = 50119
  backend_port                   = 22
  frontend_ip_configuration_name = "PublicIPAddress"
}

resource "azurerm_virtual_machine_scale_set" "kunlun_vmss" {
  name                      = "${var.env_name}-vmss"
  location                  = "${azurerm_resource_group.kunlun_resource_group.location}"
  resource_group_name       = "${azurerm_resource_group.kunlun_resource_group.name}"
  upgrade_policy_mode       = "Manual"
  network_security_group_id = "${azurerm_network_security_group.kunlun_server_network_security_group.id}"
  sku {
    name     = "${var.web_server_vm_size}"
    tier     = "Standard"
    capacity = "${var.web_server_vm_count}"
  }

  storage_profile_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }

  storage_profile_os_disk {
    name              = ""
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }

  storage_profile_data_disk {
    lun            = 0
    caching        = "ReadWrite"
    create_option  = "Empty"
    disk_size_gb   = 10
  }

  os_profile {
    computer_name_prefix = "${var.env_name}-vm"
    admin_username       = "myadmin"
    admin_password       = "Passwword1234"
  }

  os_profile_linux_config {
    disable_password_authentication = false
  }

  network_profile {
    name    = "kunlunvmssnetworkprofile"
    primary = true

    ip_configuration {
      name                                   = "kunlunvmssnetworkipconfiguration"
      primary                                = true
      subnet_id                              = "${azurerm_subnet.san_subnet.id}"
      load_balancer_backend_address_pool_ids = ["${azurerm_lb_backend_address_pool.kunlun_load_balancer_backend_address_pool.id}"]
      load_balancer_inbound_nat_rules_ids    = ["${azurerm_lb_nat_pool.kunlun_vmss_nat_pool.id}"]
    }
  }
}