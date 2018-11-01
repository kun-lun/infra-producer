package apis

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/kun-lun/infra-producer/application"
	"github.com/kun-lun/infra-producer/storage"
	"github.com/kun-lun/infra-producer/terraform"
	webserverterraform "github.com/kun-lun/infra-producer/terraform/webserver"
	artifacts "github.com/kun-lun/artifacts/pkg/apis/manifests"
)

type InfraProducer struct {
	terraformManager terraform.Manager
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
		terraformManager: terraformManager,
	}
}

func (ip InfraProducer) Setup(artifactsBytes []byte, state storage.State) error {
	im, err := artifacts.NewInfraManifestFromYAML(artifactsBytes)
	if err != nil {
		return err
	}

	ip.terraformManager.Setup(im, state)

	return nil
}

func (ip InfraProducer) Apply(state storage.State) error {
	ip.terraformManager.Apply(state)

	return nil
}
