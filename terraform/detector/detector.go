package detector

import (
	artifacts "github.com/kun-lun/artifacts/pkg/apis/manifests"
	"github.com/kun-lun/infra-producer/storage"
)

// Detector can detector whether a resource exists. If it exists,
// `DetectTemplate` will return its terraform script, and
// `DetectInput` will return related input map
type Detector interface {
	DetectTemplate(manifest artifacts.InfraManifest, state storage.State, templatesHolder interface{}) string
	DetectInput(manifest artifacts.InfraManifest, state storage.State) map[string]interface{}
}
