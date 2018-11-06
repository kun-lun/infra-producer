package handler

import (
	artifacts "github.com/kun-lun/artifacts/pkg/apis"
	"github.com/kun-lun/common/storage"
)

type TemplateGenerator interface {
	GenerateTemplate(artifacts.Manifest, storage.State) string
}
