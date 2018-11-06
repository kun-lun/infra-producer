package handler

import (
	artifacts "github.com/kun-lun/artifacts/pkg/apis"
	"github.com/kun-lun/common/storage"
)

type Manager interface {
	Setup(manifest artifacts.Manifest, kunlunState storage.State) error
	Apply(kunlunState storage.State) (storage.State, error)
}
