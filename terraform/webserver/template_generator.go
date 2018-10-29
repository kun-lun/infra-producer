package webserver

import (
	"strings"

	"github.com/xplaceholder/infra-producer/storage"
)

type templates struct {
	vars                 string
	resourceGroup        string
	network              string
	output               string
}

type TemplateGenerator struct{}

func NewTemplateGenerator() TemplateGenerator {
	return TemplateGenerator{}
}

func (t TemplateGenerator) Generate(state storage.State) string {
	tmpls := readTemplates()

	template := strings.Join([]string{tmpls.vars, tmpls.resourceGroup, tmpls.network, tmpls.output}, "\n")

	return template
}

func readTemplates() templates {
	tmpls := templates{}
	tmpls.vars = string(MustAsset("templates/vars.tf"))
	tmpls.resourceGroup = string(MustAsset("templates/resource_group.tf"))
	tmpls.network = string(MustAsset("templates/network.tf"))
	tmpls.output = string(MustAsset("templates/output.tf"))

	return tmpls
}
