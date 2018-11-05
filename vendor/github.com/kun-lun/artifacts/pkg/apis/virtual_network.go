package apis

type VirtualNetwork struct {
	Name    string   `yaml:"name"`
	Subnets []Subnet `yaml:"subnets"`
}
