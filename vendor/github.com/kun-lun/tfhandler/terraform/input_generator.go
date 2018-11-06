package terraform

import (
	"strconv"

	artifacts "github.com/kun-lun/artifacts/pkg/apis"
	"github.com/kun-lun/common/storage"
)

type InputGenerator struct {
}

func NewInputGenerator() InputGenerator {
	return InputGenerator{}
}

func (i InputGenerator) GenerateInput(manifest artifacts.Manifest, state storage.State) (map[string]interface{}, error) {
	dbUsername := manifest.Databases[0].Username
	dbPassword := manifest.Databases[0].Password
	dbStorage := strconv.Itoa(manifest.Databases[0].Storage)
	dbCore := strconv.Itoa(manifest.Databases[0].Cores)

	loadBalancerSKU := manifest.LoadBalancers[0].SKU

	vmCount := strconv.Itoa(manifest.VMGroups[0].Count)
	vmSize := manifest.VMGroups[0].SKU

	input := map[string]interface{}{
		"env_name":            state.EnvID,
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
