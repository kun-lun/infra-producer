package apis

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

// Manifest contains all needed information, all later on modules will
// use this manifest
type Manifest struct {
	Schema          string           `yaml:"schema,omitempty"`
	Region          string           `yaml:"region,omitempty"`
	IaaS            string           `yaml:"iaas,omitempty"`
	Platform        *Platform        `yaml:"platform,omitempty"`
	VMGroups        []VMGroup        `yaml:"vm_groups,omitempty"`
	VNets           []VirtualNetwork `yaml:"vnets,omitempty"`
	LoadBalancers   []LoadBalancer   `yaml:"load_balancers,omitempty"`
	StorageAccounts []StorageAccount `yaml:"storage_accounts,omitempty"`
	Databases       []Database       `yaml:"databases,omitempty"`
}

func (m *Manifest) validate() error {
	return nil
}

func (m *Manifest) autofill() error {
	for vmgIdx, vmGroup := range m.VMGroups {
		if len(vmGroup.Networks) > 0 {
			vmGroup.NetworkInfos = make([]VMNetworkInfo, len(vmGroup.Networks))
			for i, network := range vmGroup.Networks {
				if network.LoadBalancer != nil {
					vmGroup.NetworkInfos[i].LoadBalancerName = network.LoadBalancer.Name
				}
				if network.Subnet != nil {
					vmGroup.NetworkInfos[i].SubnetName = network.Subnet.Name
				}
			}
		} else if len(vmGroup.NetworkInfos) > 0 {
			vmGroup.Networks = make([]VMNetWork, len(vmGroup.NetworkInfos))
			for i, vmn := range vmGroup.NetworkInfos {
				if vmn.LoadBalancerName != "" {
					found := false
					for _, lb := range m.LoadBalancers {
						if lb.Name == vmn.LoadBalancerName {
							vmGroup.Networks[i].LoadBalancer = &lb
							found = true
						}
					}
					if !found {
						return fmt.Errorf("Can't found load balancer with given name: %s", vmn.LoadBalancerName)
					}
				}
				if vmn.SubnetName != "" {
					found := false
					for _, vnet := range m.VNets {
						for _, snet := range vnet.Subnets {
							if snet.Name == vmn.SubnetName {
								vmGroup.Networks[i].Subnet = &snet
								found = true
							}
						}
					}
					if !found {
						return fmt.Errorf("Can't found subnet with given name: %s", vmn.SubnetName)
					}
				}
			}
		}
		m.VMGroups[vmgIdx] = vmGroup
	}
	return nil
}

// ToYAML converts the object to YAML bytes
func (m *Manifest) ToYAML() (b []byte, err error) {
	if err := m.autofill(); err != nil {
		return nil, err
	}
	if err := m.validate(); err != nil {
		return nil, err
	}
	return yaml.Marshal(m)
}

// NewManifestFromYAML convert yaml bytes to Manifest object
func NewManifestFromYAML(b []byte) (m *Manifest, err error) {
	var manifest Manifest
	if err := yaml.Unmarshal(b, &manifest); err != nil {
		return nil, err
	}
	if err := manifest.autofill(); err != nil {
		return nil, err
	}
	if err := manifest.validate(); err != nil {
		return nil, err
	}

	return &manifest, nil
}
