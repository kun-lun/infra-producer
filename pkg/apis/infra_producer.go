package apis

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	artifacts "github.com/kun-lun/artifacts/pkg/apis/manifests"
	"github.com/kun-lun/common/configuration"
	"github.com/kun-lun/common/logger"
	"github.com/kun-lun/common/storage"
	"github.com/kun-lun/infra-producer/terraform"
	webserverterraform "github.com/kun-lun/infra-producer/terraform/webserver"
	"github.com/spf13/afero"
)

type InfraProducer struct {
	terraformManager terraform.Manager
}

func NewInfraProducer(globals configuration.GlobalConfiguration) InfraProducer {
	log.SetFlags(0)

	logger := logger.NewLogger(os.Stdout, os.Stdin)

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
