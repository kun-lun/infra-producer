package resources

// VMGroup contains needed information to create a set of VMs on Azure. VMs in the group
// will have the same SKU, using the same subnet.
type VMGroup struct {
	Name  string `yaml:"name"`
	Count int    `yaml:"count"`
	SKU   string `yaml:"sku"`
	Type  string `yaml:"type"`
}
