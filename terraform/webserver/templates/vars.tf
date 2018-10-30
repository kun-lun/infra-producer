variable "env_name" {}

variable "environments" {
  type = "map"
  default = {
    AzureCloud        = "public"
    AzureUSGovernment = "usgovernment"
    AzureGermanCloud  = "german"
    AzureChinaCloud   = "china"
  }
}

variable "azure_environment" {}

variable "azure_tenant_id" {}

variable "azure_subscription_id" {}

variable "azure_client_id" {}

variable "azure_client_secret" {}

variable "location" {
  default = "eastus2"
}

variable "jindou_virtual_network_address_space" {
  type    = "list"
  default = ["10.0.0.0/16"]
}
