variable "engine_version" {
  default = "5.7"
}

variable "cores" {}

variable "storage" {}

variable "backup_retention_days" {
  default = 35
}

variable "username" {}

variable "password" {}

resource "azurerm_mysql_server" "db_mysql" {
  name                      = "${var.env_name}"
  resource_group_name       = "${azurerm_resource_group.jindou_resource_group.name}"
  location                  = "${var.location}"

  sku {
    name = "GP_Gen5_${var.cores}"
    capacity = "${var.cores}"
    tier = "GeneralPurpose"
    family = "Gen5"
  }

  storage_profile {
    storage_mb = "${var.storage * 1024}"
    backup_retention_days = "${var.backup_retention_days}"
    geo_redundant_backup = "Disabled"
  }

  administrator_login = "${var.username}"
  administrator_login_password = "${var.password}"
  version = "${var.engine_version}"
  ssl_enforcement = "Enabled"
}
