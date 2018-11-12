package apis

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	artifacts "github.com/kun-lun/artifacts/pkg/apis"
	"github.com/kun-lun/common/errors"
	"github.com/kun-lun/common/logger"
	"github.com/kun-lun/common/storage"
	"github.com/kun-lun/infra-producer/handler"
	"github.com/kun-lun/tfhandler/terraform"
	"github.com/spf13/afero"
)

type InfraProducer struct {
	manager handler.Manager
}

func NewInfraProducer(stateStore storage.Store, handlerType string, debug bool) (InfraProducer, error) {
	log.SetFlags(0)

	logger := logger.NewLogger(os.Stdout, os.Stdin)

	fs := afero.NewOsFs()
	afs := &afero.Afero{Fs: fs}

	if handlerType == handler.TerraformHandlerType {
		terraformOutputBuffer := bytes.NewBuffer([]byte{})
		terraformDir, _ := stateStore.GetTerraformDir()
		dotTerraformDir := filepath.Join(terraformDir, ".terraform")
		bufferingCLI := terraform.NewCLI(terraformOutputBuffer, terraformOutputBuffer, dotTerraformDir)
		var (
			terraformCLI terraform.CLI
			out          io.Writer
		)
		if debug {
			errBuffer := io.MultiWriter(os.Stderr, terraformOutputBuffer)
			terraformCLI = terraform.NewCLI(errBuffer, terraformOutputBuffer, dotTerraformDir)
			out = os.Stdout
		} else {
			terraformCLI = bufferingCLI
			out = ioutil.Discard
		}
		terraformExecutor := terraform.NewExecutor(terraformCLI, bufferingCLI, stateStore, afs, debug, out)

		inputGenerator := terraform.NewInputGenerator()
		templateGenerator := terraform.NewTemplateGenerator()
		terraformManager := terraform.NewManager(terraformExecutor, templateGenerator, inputGenerator, terraformOutputBuffer, logger)

		return InfraProducer{
			manager: terraformManager,
		}, nil
	} else if handlerType == handler.ARMTemplateHandlerType {
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
