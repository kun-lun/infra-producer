package apis

// Database contains information to deploy a database on VM(s)
type Database struct {
	Engine               string                `yaml:"engine"`
	EngineVersion        string                `yaml:"engine_version"`
	Storage              int                   `yaml:"storage"`
	Cores                int                   `yaml:"cores"`
	BackupRetentionDays  int                   `yaml:"backup_retention_days"`
	Username             string                `yaml:"username"`
	Password             string                `yaml:"password"`
	MigrationInformation *MigrationInformation `yaml:"migrate_from,omitempty"`
}

const (
	MysqlDB = "mysql"
)

type MigrationInformation struct {
	OriginHost     string `yaml:"origin_host"`
	OriginDatabase string `yaml:"origin_database"`
	OriginUsername string `yaml:"origin_username"`
	OriginPassword string `yaml:"origin_password"`
}
