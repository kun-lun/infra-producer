package webserver

import (
	artifacts "github.com/kun-lun/artifacts/pkg/apis/manifests"
	"github.com/kun-lun/infra-producer/pkg/utils"
	"github.com/kun-lun/infra-producer/storage"
	"github.com/kun-lun/infra-producer/terraform/detector"
)

type InputGenerator struct {
	detectors []detector.Detector
}

func NewInputGenerator() InputGenerator {
	return InputGenerator{
		detectors: []detector.Detector{
			newInfraDetector(),
			newLoadBalancerDetector(),
		},
	}
}

func (i InputGenerator) Generate(manifest artifacts.InfraManifest, state storage.State) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for _, d := range i.detectors {
		result = utils.MergeMap(
			result,
			d.DetectInput(manifest, state),
		)
	}
	return result, nil
}

func (i InputGenerator) Credentials(state storage.State) map[string]string {
	return map[string]string{
		"azure_environment":     state.Azure.Environment,
		"azure_subscription_id": state.Azure.SubscriptionID,
		"azure_tenant_id":       state.Azure.TenantID,
		"azure_client_id":       state.Azure.ClientID,
		"azure_client_secret":   state.Azure.ClientSecret,
	}
}
