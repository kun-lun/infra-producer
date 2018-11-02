package webserver

import (
	artifacts "github.com/kun-lun/artifacts/pkg/apis/manifests"
	"github.com/kun-lun/infra-producer/storage"
	"github.com/kun-lun/infra-producer/terraform/detector"
)

type templates struct {
	mysql         string
	jumpbox       string
	loadBalancer  string
	output        string
	provider      string
	resourceGroup string
	subnet        string
	vars          string
	vmServer      string
	vmssServer    string
	vnet          string
}

type TemplateGenerator struct {
	templates templates
	detectors []detector.Detector
}

func NewTemplateGenerator() TemplateGenerator {
	tg := TemplateGenerator{
		templates: readTemplates(),
	}
	tg.detectors = []detector.Detector{
		newInfraDetector(),
		newLoadBalancerDetector(),
	}
	return tg
}

func (t TemplateGenerator) Generate(manifest artifacts.InfraManifest, state storage.State) string {
	var template string
	for _, d := range t.detectors {
		template += d.DetectTemplate(manifest, state, t.templates)
	}
	return template
}

func readTemplates() templates {
	tmpls := templates{}
	tmpls.mysql = string(MustAsset("templates/db_mysql.tf"))
	tmpls.jumpbox = string(MustAsset("templates/jumpbox.tf"))
	tmpls.loadBalancer = string(MustAsset("templates/load_balancer.tf"))
	tmpls.output = string(MustAsset("templates/output.tf"))
	tmpls.provider = string(MustAsset("templates/provider.tf"))
	tmpls.resourceGroup = string(MustAsset("templates/resource_group.tf"))
	tmpls.subnet = string(MustAsset("templates/subnet.tf"))
	tmpls.vars = string(MustAsset("templates/vars.tf"))
	tmpls.vmServer = string(MustAsset("templates/vm_web_server.tf"))
	tmpls.vmssServer = string(MustAsset("templates/vmss_web_server.tf"))
	tmpls.vnet = string(MustAsset("templates/vnet.tf"))
	return tmpls
}
