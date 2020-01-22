package config

// Server holds data server configuration data
type Server struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}
