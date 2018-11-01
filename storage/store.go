package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/kun-lun/common/fileio"
)

var (
	marshalIndent = json.MarshalIndent
	uuidNewV4     = uuid.NewV4
)

const (
	STATE_SCHEMA = 1
	STATE_FILE   = "kunlun-state.json"
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

	state.Version = s.stateSchema

	if state.ID == "" {
		uuid, err := uuidNewV4()
		if err != nil {
			return fmt.Errorf("Create state ID: %s", err)
		}
		state.ID = uuid.String()
	}

	jsonData, err := marshalIndent(state, "", "\t")
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

func (s Store) GetTerraformDir() (string, error) {
	return s.getDir("terraform", os.ModePerm)
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
