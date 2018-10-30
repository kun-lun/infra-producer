package manifests

import (
	"github.com/xplaceholder/artifacts/pkg/apis/basic"
	"github.com/xplaceholder/artifacts/pkg/apis/resources"
	yaml "gopkg.in/yaml.v2"
)

// InfraManifest contains all needed information about the infrastructure,
// and will be consumed by infra module
type InfraManifest struct {
	Schema       string                  `yaml:"schema,omitempty"`
	Platform     *basic.Platform         `yaml:"platform,omitempty"`
	Region       string                  `yaml:"region,omitempty"`
	LoadBalancer *resources.LoadBalancer `yaml:"load_balancer,omitempty"`
	VMGroups     []resources.VMGroup     `yaml:"vm_groups,omitempty"`
}

// ToYAML converts the object to YAML bytes
func (im InfraManifest) ToYAML() (b []byte, err error) {
	return yaml.Marshal(im)
}

// NewInfraManifestFromYAML convert yaml bytes to InfraManifest object
func NewInfraManifestFromYAML(b []byte) (m InfraManifest, err error) {
	var im InfraManifest
	err = yaml.Unmarshal(b, &im)
	return im, err
}
