package config

type Config struct {
	Address     string
	Port        string
	SSLCertPath string
	SSLKeyPath  string
	JWTSecret   string
	JWTExp      string
}
