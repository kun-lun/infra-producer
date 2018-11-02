package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/kun-lun/common/fileio"
	uuid "github.com/nu7hatch/gouuid"
)

const (
	STATE_SCHEMA = 1
	STATE_FILE   = "kl-state.json"
)

type Store struct {
	dir         string
	fs          fs
	stateSchema int
}

type fs interface {
	fileio.FileWriter
	fileio.Remover
	fileio.AllRemover
	fileio.Stater
	fileio.AllMkdirer
	fileio.DirReader
}

func NewStore(dir string, fs fs) Store {
	return Store{
		dir:         dir,
		fs:          fs,
		stateSchema: STATE_SCHEMA,
	}
}

func (s Store) Set(state State) error {
	_, err := s.fs.Stat(s.dir)
	if err != nil {
		return fmt.Errorf("Stat state dir: %s", err)
	}

	if reflect.DeepEqual(state, State{}) {
		return nil
	}

	state.Version = s.stateSchema

	if state.ID == "" {
		uuid, err := uuid.NewV4()
		if err != nil {
			return fmt.Errorf("Create state ID: %s", err)
		}
		state.ID = uuid.String()
	}

	jsonData, err := json.MarshalIndent(state, "", "\t")
	if err != nil {
		return err
	}

	stateFile := filepath.Join(s.dir, STATE_FILE)
	err = s.fs.WriteFile(stateFile, jsonData, os.FileMode(0644))
	if err != nil {
		return err
	}

	return nil
}

func (s Store) GetStateDir() string {
	return s.dir
}

// GetInfraDir get the infrastructure folder.
func (s Store) GetInfraDir() (string, error) {
	return s.getDir("infra", os.ModePerm)
}

// GetTerraformDir get the terraform folder, this should be the sub folder of infra.
func (s Store) GetTerraformDir() (string, error) {
	return s.getDir("infra/terraform", os.ModePerm)
}

// GetArtifactsDir get artifacts folder
func (s Store) GetArtifactsDir() (string, error) {
	return s.getDir("artifacts", os.ModePerm)
}

func (s Store) GetDeploymentsDir() (string, error) {
	return s.getDir("deployments", os.ModePerm)
}

func (s Store) GetAnsibleDir() (string, error) {
	return s.getDir("deployments/ansible", os.ModePerm)
}

func (s Store) GetVarsDir() (string, error) {
	return s.getDir("vars", StateMode)
}

func (s Store) getDir(name string, perm os.FileMode) (string, error) {
	dir := filepath.Join(s.dir, name)
	err := s.fs.MkdirAll(dir, perm)
	if err != nil {
		return "", fmt.Errorf("Get %s dir: %s", name, err)
	}
	return dir, nil
}
