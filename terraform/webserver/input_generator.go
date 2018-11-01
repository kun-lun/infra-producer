package webserver

import (
	"strconv"

	artifacts "github.com/kun-lun/artifacts/pkg/apis/manifests"
	"github.com/kun-lun/infra-producer/storage"
)

type InputGenerator struct {
}

func NewInputGenerator() InputGenerator {
	return InputGenerator{}
}

func (i InputGenerator) Generate(manifest artifacts.InfraManifest, state storage.State) (map[string]interface{}, error) {
	dbUsername := manifest.Database.Username
	dbPassword := manifest.Database.Password
	dbStorage := strconv.Itoa(manifest.Database.Storage)
	dbCore := strconv.Itoa(manifest.Database.Cores)

	loadBalancerSKU := manifest.LoadBalancer.SKU

	vmCount := strconv.Itoa(manifest.VMGroups[0].Count)
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
