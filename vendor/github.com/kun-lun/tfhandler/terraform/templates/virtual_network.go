package templates

import artifacts "github.com/kun-lun/artifacts/pkg/apis"

var virtualNetworkTF = []byte(`
resource "azurerm_virtual_network" "{{.vnetName}}" {
	name                = "${var.env_name}-{{.vnetName}}"
	resource_group_name = "${azurerm_resource_group.kunlun_resource_group.name}"
	address_space       = ["10.0.0.0/16"]
	location            = "${var.location}"
}
`)

var subnetTF = []byte(`
resource "azurerm_subnet" "{{.subnetName}}" {
	name                 = "${var.env_name}-{{.subnetName}}"
	resource_group_name  = "${azurerm_resource_group.kunlun_resource_group.name}"
	virtual_network_name = "${azurerm_virtual_network.{{.vnetName}}.name}"
	address_prefix       = "${var.{{.subnetName}}_as_address_prefix}"
}

variable "{{.subnetName}}_as_address_prefix" {}
`)

var subnetTFVars = []byte(`
{{.subnetName}}_as_address_prefix = "{{.as_address_prefix}}"
`)

func NewVirtualNetworkTemplate(vnet artifacts.VirtualNetwork) (string, error) {
	tf := ""

	vnetTF, err := render(virtualNetworkTF, getVirtualNetworkTFParams(vnet))
	if err != nil {
		return "", err
	}
	tf += vnetTF

	for _, snet := range vnet.Subnets {
		snetTF, err := render(subnetTF, getSubnetTFParams(snet, vnet.Name))
		if err != nil {
			return "", err
		}
		tf += snetTF
	}

	return tf, nil
}

func NewVirtualNetworkInput(vnet artifacts.VirtualNetwork) (string, error) {
	tfvars := ""
	for _, snet := range vnet.Subnets {
		snetTFVars, err := render(subnetTFVars, getSubnetTFVarsParams(snet))
		if err != nil {
			return "", err
		}
		tfvars += snetTFVars
	}
	return tfvars, nil
}

func getVirtualNetworkTFParams(vnet artifacts.VirtualNetwork) map[string]interface{} {
	return map[string]interface{}{
		"vnetName": vnet.Name,
	}
}

func getSubnetTFParams(snet artifacts.Subnet, vnetName string) map[string]interface{} {
	return map[string]interface{}{
		"subnetName": snet.Name,
		"vnetName":   vnetName,
	}
}

func getSubnetTFVarsParams(snet artifacts.Subnet) map[string]interface{} {
	return map[string]interface{}{
		"subnetName":        snet.Name,
		"as_address_prefix": snet.Range,
	}
}
