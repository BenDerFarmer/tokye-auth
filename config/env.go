package config

import (
	"os"
	"strconv"
	"strings"
)

type SecretConfig struct {
	RefreshToken string
	SqlDSN       string
}

type SmtpConfig struct {
	Domain   string
	Port     int
	Username string
	Password string
	FromMail string
}

type PassKeyConfig struct {
	DisplayName string
	RPID        string
	Origins     []string
}

type RedisConfig struct {
	Address  string
	Password string
	Username string
	Database int
}

type Config struct {
	Secret     SecretConfig
	Redis      RedisConfig
	Mail       SmtpConfig
	PassKey    PassKeyConfig
	CORSOrigns []string
	DebugMode  bool
	Port       string
}

func New() *Config {
	return &Config{
		Secret: SecretConfig{
			RefreshToken: getEnv("REFRESH_TOKEN_SECRET", ""),
			SqlDSN:       getEnv("SQL_DSN", ""),
		},
		Mail: SmtpConfig{
			Domain:   getEnv("SMTP_DOMAIN", ""),
			Port:     getEnvAsInt("SMTP_PORT", 465),
			Username: getEnv("SMTP_USERNAME", ""),
			Password: getEnv("SMTP_PASSWORD", ""),
			FromMail: getEnv("SMTP_FROM_MAIL", getEnv("SMTP_USERNAME", "")),
		},
		Redis: RedisConfig{
			Address:  getEnv("REDIS_ADDRESS", "localhost:6379"),
			Username: getEnv("REDIS_USERNAME", ""),
			Password: getEnv("REDIS_PASSWORD", ""),
			Database: getEnvAsInt("REDIS_DATABASE", 0),
		},
		PassKey: PassKeyConfig{
			DisplayName: getEnv("PASSKEY_DISPLAYNAME", "Auth"),
			RPID:        getEnv("PASSKEY_PRID", ""),
			Origins:     getEnvAsSlice("PASSKEY_ORIGINS", []string{}, ";"),
		},
		CORSOrigns: getEnvAsSlice("CORS_ORIGINS", []string{}, ";"),
		DebugMode:  getEnvAsBool("DEBUG_MODE", false),
		Port:       getEnv("PORT", "3000"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
