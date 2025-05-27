package model

type Config struct {
	PostgresDbConfig *PgDbConfig `json:"postgres"`
	// Add other configurations as needed
}

type PgDbConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}
