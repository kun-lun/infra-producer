package resources

// LoadBalancer contains needed information to create a load balancer on Azure.
type LoadBalancer struct {
	SKU string `yaml:"sku"`
}
