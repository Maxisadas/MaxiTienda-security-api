package configuration

type Config struct {
	Database_host     string
	Database_user     string
	Database_password string
	Database_dbname   string
	Database_port     int
	Database_sslmode  string
}

func GetConfig() *Config {
	var config = new(Config)
	config.Database_host = "localhost"
	config.Database_dbname = "tienda"
	config.Database_user = "postgres"
	config.Database_password = "postgres"
	config.Database_port = 5432
	config.Database_sslmode = "disable"
	return config
}
