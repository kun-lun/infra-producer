package webserver

import (
	"strconv"
	"strings"

	artifacts "github.com/kun-lun/artifacts/pkg/apis/manifests"
	"github.com/kun-lun/infra-producer/storage"
)

type infraDetector struct{}

func (id infraDetector) DetectTemplate(
	manifest artifacts.InfraManifest,
	state storage.State,
	templatesHolder interface{},
) string {
	tmpls := templatesHolder.(templates)
	template := strings.Join(
		[]string{
			tmpls.mysql,
			tmpls.jumpbox,
			tmpls.output,
			tmpls.provider,
			tmpls.resourceGroup,
			tmpls.subnet,
			tmpls.vars,
			tmpls.vmssServer,
			tmpls.vnet,
		},
		"\n",
	)

	return template
}

func (id infraDetector) DetectInput(
	manifest artifacts.InfraManifest,
	state storage.State,
) map[string]interface{} {
	dbUsername := manifest.Database.Username
	dbPassword := manifest.Database.Password
	dbStorage := strconv.Itoa(manifest.Database.Storage)
	dbCore := strconv.Itoa(manifest.Database.Cores)

	vmCount := strconv.Itoa(manifest.VMGroups[0].Count)
	vmSize := manifest.VMGroups[0].SKU

	result := map[string]interface{}{
		"env_name":            state.EnvName,
		"region":              state.Azure.Region,
		"database_username":   dbUsername,
		"database_password":   dbPassword,
		"storage":             dbStorage,
		"cores":               dbCore,
		"web_server_vm_count": vmCount,
		"web_server_vm_size":  vmSize,
	}
	return result
}

func newInfraDetector() infraDetector {
	return infraDetector{}
}
