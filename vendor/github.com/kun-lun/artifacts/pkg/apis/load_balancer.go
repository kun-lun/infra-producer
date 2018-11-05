package apis

// LoadBalancer contains needed information to create a load balancer on Azure.
type LoadBalancer struct {
	Name string `yaml:"name"`
	SKU  string `yaml:"sku"`
}

const (
	LoadBalancerStandardSKU = "standard"
	LoadBalancerBasicSKU    = "basic"
)
