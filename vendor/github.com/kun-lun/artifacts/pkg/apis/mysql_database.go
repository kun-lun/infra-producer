package apis

// Database contains information to deploy a database on VM(s)
type MysqlDatabase struct {
	Name                 string                `yaml:"name"`
	Version              string                `yaml:"version"`
	Cores                int                   `yaml:"cores"`
	Tier                 string                `yaml:"tier"`
	Family               string                `yaml:"family"`
	Storage              int                   `yaml:"storage"`
	BackupRetentionDays  int                   `yaml:"backup_retention_days"`
	SSLEnforcement       string                `yaml:"ssl_enforcement"`
	Username             string                `yaml:"username"`
	Password             string                `yaml:"password"`
	MigrationInformation *MigrationInformation `yaml:"migrate_from,omitempty"`
}

type MigrationInformation struct {
	OriginHost     string `yaml:"origin_host"`
	OriginDatabase string `yaml:"origin_database"`
	OriginUsername string `yaml:"origin_username"`
	OriginPassword string `yaml:"origin_password"`
}
