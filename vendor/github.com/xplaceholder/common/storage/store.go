package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/xplaceholder/common/fileio"
)

const (
	STATE_SCHEMA = 14
	STATE_FILE   = "kid-state.json"
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
