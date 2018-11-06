package apis

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	artifacts "github.com/kun-lun/artifacts/pkg/apis"
	"github.com/kun-lun/common/configuration"
	"github.com/kun-lun/common/errors"
	"github.com/kun-lun/common/logger"
	"github.com/kun-lun/common/storage"
	"github.com/kun-lun/tfhandler/terraform"
	"github.com/kun-lun/infra-producer/handler"
	"github.com/spf13/afero"
)

type InfraProducer struct {
	manager handler.Manager
}

func NewInfraProducer(globals configuration.GlobalConfiguration) (InfraProducer, error) {
	log.SetFlags(0)

	logger := logger.NewLogger(os.Stdout, os.Stdin)

	fs := afero.NewOsFs()
	afs := &afero.Afero{Fs: fs}

	stateStore := storage.NewStore(globals.StateDir, afs)

	if globals.HandlerType == handler.TerraformHandlerType {
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

		inputGenerator := terraform.NewInputGenerator()
		templateGenerator := terraform.NewTemplateGenerator()
		terraformManager := terraform.NewManager(terraformExecutor, templateGenerator, inputGenerator, terraformOutputBuffer, logger)

		return InfraProducer{
			manager: terraformManager,
		}, nil
	} else if globals.HandlerType == handler.ARMTemplateHandlerType {
		return InfraProducer{}, &errors.NotImplementedError{}
	} else {
		return InfraProducer{}, &errors.NotSupportedError{}
	}
}

func (ip InfraProducer) Setup(artifactsBytes []byte, state storage.State) error {
	im, err := artifacts.NewManifestFromYAML(artifactsBytes)
	if err != nil {
		return err
	}

	ip.manager.Setup(*im, state)

	return nil
}

func (ip InfraProducer) Apply(state storage.State) error {
	ip.manager.Apply(state)

	return nil
}
