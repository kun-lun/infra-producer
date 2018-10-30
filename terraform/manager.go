package terraform

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/xplaceholder/infra-producer/storage"
	artifacts "github.com/xplaceholder/artifacts/pkg/apis/manifests"
	"github.com/coreos/go-semver/semver"
)

type Manager struct {
	executor              executor
	templateGenerator     TemplateGenerator
	inputGenerator        InputGenerator
	terraformOutputBuffer *bytes.Buffer
	logger                logger
}

type executor interface {
	Version() (string, error)
	Setup(terraformTemplate string, inputs map[string]interface{}) error
	Init() error
	Apply(credentials map[string]string) error
	Validate(credentials map[string]string) error
	Destroy(credentials map[string]string) error
	Outputs() (map[string]interface{}, error)
	Output(string) (string, error)
	IsPaved() (bool, error)
}

type InputGenerator interface {
	Generate(artifacts.InfraManifest, storage.State) (map[string]interface{}, error)
	Credentials(state storage.State) map[string]string
}

type TemplateGenerator interface {
	Generate(artifacts.InfraManifest, storage.State) string
}

type logger interface {
	Step(string, ...interface{})
}

func NewManager(executor executor, templateGenerator TemplateGenerator, inputGenerator InputGenerator, terraformOutputBuffer *bytes.Buffer, logger logger) Manager {
	return Manager{
		executor:              executor,
		templateGenerator:     templateGenerator,
		inputGenerator:        inputGenerator,
		terraformOutputBuffer: terraformOutputBuffer,
		logger:                logger,
	}
}

func (m Manager) Version() (string, error) {
	return m.executor.Version()
}

func (m Manager) ValidateVersion() error {
	version, err := m.executor.Version()
	if err != nil {
		return err
	}

	currentVersion, err := semver.NewVersion(version)
	if err != nil {
		return err
	}

	minimumVersion, err := semver.NewVersion("0.11.0")
	if err != nil {
		return err
	}

	if currentVersion.LessThan(*minimumVersion) {
		return errors.New("Terraform version must be at least v0.11.0")
	}

	return nil
}

func (m Manager) Setup(manifest artifacts.InfraManifest, jindouState storage.State) error {
	m.logger.Step("generating terraform template")
	template := m.templateGenerator.Generate(manifest, jindouState)

	m.logger.Step("generating terraform variables")
	input, err := m.inputGenerator.Generate(manifest, jindouState)
	if err != nil {
		return fmt.Errorf("Input generator generate: %s", err)
	}

	if err := m.executor.Setup(template, input); err != nil {
		return fmt.Errorf("Executor setup: %s", err)
	}

	return m.Init(jindouState)
}

func (m Manager) Init(jindouState storage.State) error {
	m.logger.Step("terraform init")
	if err := m.executor.Init(); err != nil {
		return fmt.Errorf("Executor init: %s", err)
	}
	return nil
}

func (m Manager) Apply(jindouState storage.State) (storage.State, error) {
	m.logger.Step("terraform init")
	if err := m.executor.Init(); err != nil {
		return jindouState, fmt.Errorf("Executor init: %s", err)
	}

	m.logger.Step("terraform apply")
	err := m.executor.Apply(m.inputGenerator.Credentials(jindouState))

	jindouState.LatestTFOutput = readAndReset(m.terraformOutputBuffer)

	if err != nil {
		return jindouState, fmt.Errorf("Executor apply: %s", err)
	}

	return jindouState, nil
}

func (m Manager) Destroy(jindouState storage.State) (storage.State, error) {
	m.logger.Step("terraform destroy")
	err := m.executor.Destroy(m.inputGenerator.Credentials(jindouState))

	jindouState.LatestTFOutput = readAndReset(m.terraformOutputBuffer)

	if err != nil {
		return jindouState, fmt.Errorf("Executor destroy: %s", err)
	}

	m.logger.Step("finished destroying infrastructure")
	return jindouState, nil
}

func (m Manager) Validate(jindouState storage.State) (storage.State, error) {
	m.logger.Step("terraform validate")
	err := m.executor.Validate(m.inputGenerator.Credentials(jindouState))

	jindouState.LatestTFOutput = readAndReset(m.terraformOutputBuffer)

	if err != nil {
		return jindouState, fmt.Errorf("Executor validate: %s", err)
	}

	return jindouState, nil
}

func (m Manager) GetOutputs() (Outputs, error) {
	tfOutputs, err := m.executor.Outputs()
	if err != nil {
		return Outputs{}, err
	}

	return Outputs{Map: tfOutputs}, nil
}

func (m Manager) IsPaved() (bool, error) {
	return m.executor.IsPaved()
}

func readAndReset(buf *bytes.Buffer) string {
	contents := buf.Bytes()
	buf.Reset()

	return string(contents)
}
