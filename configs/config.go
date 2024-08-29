package configs

import "os"

type Config struct {
	DBHost	string
	DBUser	string
	DBPass	string
	DBName	string
	DBPort	string

	JWTKey	string
}
func New() *Config {
	return &Config{
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		DBPort: os.Getenv("DB_PORT"),
		
		JWTKey: os.Getenv("JWT_SECRET"),
	}
}
