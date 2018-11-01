package webserver

import (
	artifacts "github.com/xplaceholder/artifacts/pkg/apis/manifests"
	"github.com/xplaceholder/infra-producer/storage"
)

type InputGenerator struct {
}

func NewInputGenerator() InputGenerator {
	return InputGenerator{}
}

func (i InputGenerator) Generate(manifest artifacts.InfraManifest, state storage.State) (map[string]interface{}, error) {
	dbUsername := manifest.Database.Username
	dbPassword := manifest.Database.Password
	dbStorage := manifest.Database.Storage
	dbCore := manifest.Database.Cores

	loadBalancerSKU := manifest.LoadBalancer.SKU

	vmCount := manifest.VMGroups[0].Count
	vmSize := manifest.VMGroups[0].SKU

	input := map[string]interface{}{
		"env_name":            state.EnvName,
		"region":              state.Azure.Region,
		"database_username":   dbUsername,
		"database_password":   dbPassword,
		"storage":             dbStorage,
		"cores":               dbCore,
		"web_server_vm_count": vmCount,
		"web_server_vm_size":  vmSize,
		"load_balancer_sku":   loadBalancerSKU,
	}

	return input, nil
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
