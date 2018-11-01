package manifests

import (
	"github.com/xplaceholder/artifacts/pkg/apis/resources"
	yaml "gopkg.in/yaml.v2"
)

// DeployManifest contains all needed information about the deployment,
// and will be consumed by deploy module
type DeployManifest struct {
	Databases []resources.Database `yaml:"databases,omitempty"`
}

// ToYAML converts the object to YAML bytes
func (dm DeployManifest) ToYAML() (b []byte, err error) {
	return yaml.Marshal(dm)
}

// NewDeployManifestFromYAML convert yaml bytes to DeployManifest object
func NewDeployManifestFromYAML(b []byte) (dm DeployManifest, err error) {
	var im DeployManifest
	err = yaml.Unmarshal(b, &im)
	return im, err
}
