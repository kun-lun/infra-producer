package resources

// Database contains information to deploy a database on VM(s)
type Database struct {
	Engine              string `yaml:"engine"`
	EngineVersion       string `yaml:"engine_version"`
	Storage             int    `yaml:"storage"`
	Cores               int    `yaml:"cores"`
	BackupRetentionDays int    `yaml:"backup_retention_days"`
	Username            string `yaml:"username"`
	Password            string `yaml:"password"`
}

const (
	MysqlDB = "mysql"
)
