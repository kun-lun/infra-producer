resource "azurerm_public_ip" "jindou_load_balancer_ip" {
  name                         = "${var.env_name}-load-balancer-ip"
  location                     = "${azurerm_resource_group.jindou_resource_group.location}"
  resource_group_name          = "${azurerm_resource_group.jindou_resource_group.name}"
  public_ip_address_allocation = "static"
  sku                          = "${var.load_balancer_sku}"
}

resource "azurerm_lb" "jindou_load_balancer" {
  name                = "${var.env_name}-lb"
  location            = "${azurerm_resource_group.jindou_resource_group.location}"
  resource_group_name = "${azurerm_resource_group.jindou_resource_group.name}"
  sku                 = "${var.load_balancer_sku}"
  frontend_ip_configuration {
    name                 = "PublicIPAddress"
    public_ip_address_id = "${azurerm_public_ip.jindou_load_balancer_ip.id}"
  }
}

resource "azurerm_lb_backend_address_pool" "jindou_load_balancer_backend_address_pool" {
  resource_group_name = "${azurerm_resource_group.jindou_resource_group.name}"
  loadbalancer_id     = "${azurerm_lb.jindou_load_balancer.id}"
  name                = "BackEndAddressPool"
}

resource "azurerm_lb_probe" "jindou_load_balancer_probe" {
  resource_group_name = "${azurerm_resource_group.jindou_resource_group.name}"
  loadbalancer_id     = "${azurerm_lb.jindou_load_balancer.id}"
  name                = "http-serving-probe"
  protocol            = "Http"
  request_path        = "/"
  port                = 80
}

resource "azurerm_lb_rule" "jondou_load_balancer_rule" {
  resource_group_name            = "${azurerm_resource_group.jindou_resource_group.name}"
  loadbalancer_id                = "${azurerm_lb.jindou_load_balancer.id}"
  name                           = "LBRule"
  protocol                       = "Tcp"
  frontend_port                  = 80
  backend_port                   = 80
  frontend_ip_configuration_name = "PublicIPAddress"
  backend_address_pool_id        = "${azurerm_lb_backend_address_pool.jindou_load_balancer_backend_address_pool.id}"
  probe_id                       = "${azurerm_lb_probe.jindou_load_balancer_probe.id}"
}