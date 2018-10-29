package apis

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/xplaceholder/infra-producer/application"
	"github.com/xplaceholder/infra-producer/storage"
	"github.com/xplaceholder/infra-producer/terraform"
	webserverterraform "github.com/xplaceholder/infra-producer/terraform/webserver"
	artifacts "github.com/xplaceholder/artifacts/pkg/api/manifests"
)

type InfraProducer struct {
	TerraformManager terraform.Manager
}

func NewInfraProducer(globals application.GlobalConfiguration) InfraProducer {
	log.SetFlags(0)

	logger := application.NewLogger(os.Stdout, os.Stdin)

	// File IO
	fs := afero.NewOsFs()
	afs := &afero.Afero{Fs: fs}

	// bbl Configuration
	stateStore := storage.NewStore(globals.StateDir, afs)

	terraformOutputBuffer := bytes.NewBuffer([]byte{})
	dotTerraformDir := filepath.Join(globals.StateDir, "terraform", ".terraform")
	bufferingCLI := terraform.NewCLI(terraformOutputBuffer, terraformOutputBuffer, dotTerraformDir)
	var (
		terraformCLI terraform.CLI
		out          io.Writer
	)
	if globals.Debug {
		errBuffer := io.MultiWriter(os.Stderr, terraformOutputBuffer)
		terraformCLI = terraform.NewCLI(errBuffer, terraformOutputBuffer, dotTerraformDir)
		out = os.Stdout
	} else {
		terraformCLI = bufferingCLI
		out = ioutil.Discard
	}
	terraformExecutor := terraform.NewExecutor(terraformCLI, bufferingCLI, stateStore, afs, globals.Debug, out)

	inputGenerator := webserverterraform.NewInputGenerator()
	templateGenerator := webserverterraform.NewTemplateGenerator()
	terraformManager := terraform.NewManager(terraformExecutor, templateGenerator, inputGenerator, terraformOutputBuffer, logger)

	return InfraProducer{
		TerraformManager: terraformManager,
	}
}

func (ip InfraProducer) GetRegion(b []byte) string {
	im, err := artifacts.NewInfraManifestFromYAML(b)
	if err != nil {
		return ""
	}
	return im.Region
}
