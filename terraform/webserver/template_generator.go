package webserver

import (
	"strings"

	artifacts "github.com/xplaceholder/artifacts/pkg/apis/manifests"
	"github.com/xplaceholder/infra-producer/storage"
)

type templates struct {
	provider      string
	vars          string
	resourceGroup string
	network       string
	vmss          string
	loadBalancer  string
	devbox        string
	output        string
	nsg           string
}

type TemplateGenerator struct{}

func NewTemplateGenerator() TemplateGenerator {
	return TemplateGenerator{}
}

func (t TemplateGenerator) Generate(manifest artifacts.InfraManifest, state storage.State) string {
	tmpls := readTemplates()

	template := strings.Join(
		[]string{tmpls.provider, tmpls.vars, tmpls.resourceGroup, tmpls.network, tmpls.output, tmpls.vmss, tmpls.devbox, tmpls.loadBalancer, tmpls.nsg},
		"\n",
	)

	return template
}

func readTemplates() templates {
	tmpls := templates{}
	tmpls.provider = string(MustAsset("templates/provider.tf"))
	tmpls.vars = string(MustAsset("templates/vars.tf"))
	tmpls.resourceGroup = string(MustAsset("templates/resource_group.tf"))
	tmpls.network = string(MustAsset("templates/network.tf"))
	tmpls.devbox = string(MustAsset("templates/devbox.tf"))
	tmpls.vmss = string(MustAsset("templates/vmss.tf"))
	tmpls.loadBalancer = string(MustAsset("templates/load_balancer.tf"))
	tmpls.output = string(MustAsset("templates/output.tf"))
	tmpls.nsg = string(MustAsset("templates/network_security_group.tf"))
	return tmpls
}
