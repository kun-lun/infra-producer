package apis

// LoadBalancer contains needed information to create a load balancer on Azure.
type LoadBalancer struct {
	Name                string                           `yaml:"name"`
	SKU                 string                           `yaml:"sku"`
	BackendAddressPools []LoadBalancerBackendAddressPool `yaml:"backend_address_pools"`
	HealthProbes        []LoadBalancerHealthProbe        `yaml:"health_probes"`
	Rules               []LoadBalancerRule               `yaml:"rules"`
}

type LoadBalancerBackendAddressPool struct {
	Name string `yaml:"name"`
}

type LoadBalancerHealthProbe struct {
	Name        string `yaml:"name"`
	Protocol    string `yaml:"protocol"`
	Port        int    `yaml:"port"`
	RequestPath string `yaml:"request_path"`
}

type LoadBalancerRule struct {
	Name                   string `yaml:"name"`
	Protocol               string `yaml:"protocol"`
	FrontendPort           int    `yaml:"frontend_port"`
	BackendPort            int    `yaml:"backend_port"`
	BackendAddressPoolName string `yaml:"backend_address_pool_name"`
	HealthProbeName        string `yaml:"health_probe_name"`
}

const (
	LoadBalancerStandardSKU = "standard"
	LoadBalancerBasicSKU    = "basic"
)
