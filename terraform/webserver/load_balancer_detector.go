package webserver

import (
	artifacts "github.com/kun-lun/artifacts/pkg/apis/manifests"
	"github.com/kun-lun/infra-producer/storage"
)

type loadBalancerDetector struct{}

func (lbd loadBalancerDetector) DetectTemplate(
	manifest artifacts.InfraManifest,
	state storage.State,
	templatesHolder interface{},
) string {
	lb := manifest.LoadBalancer
	if lb == nil {
		return ""
	}

	templates := templatesHolder.(templates)
	return templates.loadBalancer + "\n"
}

func (lbd loadBalancerDetector) DetectInput(
	manifest artifacts.InfraManifest,
	state storage.State,
) map[string]interface{} {
	lb := manifest.LoadBalancer
	result := make(map[string]interface{})

	if lb == nil {
		return result
	}
	result["load_balancer_sku"] = lb.SKU
	return result
}

func newLoadBalancerDetector() loadBalancerDetector {
	return loadBalancerDetector{}
}
