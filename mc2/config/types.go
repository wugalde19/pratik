package config

// Server holds data server configuration data
type Server struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

// DB holds database configuration data
type DB struct {
  Host string `yaml:"host"`
  Port int `yaml:"port"`
  User string `yaml:"user"`
  Password string `yaml:"password"`
  DBName string `yaml:"db_name"`
}

// JWT holds JWT configuration data
type JWT struct {
	Duration        	int    `yaml:"duration_minutes"`
	RefreshDuration 	int    `yaml:"refresh_duration_minutes"`
	SigningAlgorithm	string `yaml:"signing_algorithm"`
	SigningKeyEnv 		string `yaml:"signing_key_env"`
}
