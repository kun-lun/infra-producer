package storage

type State struct {
	Version        int       `json:"version"`
	ID             string    `json:"id"`
	EnvName        string    `json:"envName"`
	Azure          Azure     `json:"azure,omitempty"`
	LatestTFOutput string    `json:"latestTFOutput"`
}
