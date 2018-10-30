package resources

// Database contains information to deploy a database on VM(s)
type Database struct {
	Type    string `yaml:"type"`
	Version string `yaml:"version"`
	// Storage capacity, in MB
	Storage string `yaml:"storage"`
}

const (
	MysqlDB = "mysql"
)
