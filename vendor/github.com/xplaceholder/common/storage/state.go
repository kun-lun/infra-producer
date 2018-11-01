package storage

type State struct {
	Version    int    `json:"version"`
	KIDVersion string `json:"kidVersion"`
	IAAS       string `json:"iaas"`
	ID         string `json:"id"`
	EnvID      string `json:"envID"`
}
