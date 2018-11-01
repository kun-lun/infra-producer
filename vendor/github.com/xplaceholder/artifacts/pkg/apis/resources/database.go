package resources

// Database contains information to deploy a database on VM(s)
type Database struct {
	Type    string `yaml:"type"`
	Version string `yaml:"version"`
	// Storage capacity, in MB
	Storage        string `yaml:"storage"`
	OriginHost     string `yaml:"origin_host"`
	OriginDatabase string `yaml:"origin_database"`
	OriginUsername string `yaml:"origin_username"`
	OriginPassword string `yaml:"origin_password"`
	EnvVarHost     string `yaml:"env_var_host"`
	EnvVarDatabase string `yaml:"env_var_database"`
	EnvVarUsername string `yaml:"env_var_username"`
	EnvVarPassword string `yaml:"env_var_password"`
}

const (
	MysqlDB = "mysql"
)
