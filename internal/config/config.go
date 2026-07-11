package config

import (
	"io"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	DB     DbConfig
	JWT    JWTConfig
	Server ServerConfig
	Logger LoggerConfig
}

type ServerConfig struct {
	Port string `env:"PORT"`
}

type DbConfig struct {
	Path         string `env:"DB_PATH"`
	MaxOpenConns int    `env:"DB_MAX_OPEN_CONNS"`
	InitPath     string `env:"DB_INIT_PATH"`
}

type JWTConfig struct {
	Secret   string        `env:"JWT_SECRET"`
	ExpTime  time.Duration `env:"JWT_EXP_TIME"`
	Audience string        `env:"JWT_AUDIENCE"`
	Issuer   string        `env:"JWT_ISSUER"`
}

type LoggerConfig struct {
	FilePath  string `env:"LOG_FILE"`
	MaxSize   int    `env:"LOG_MAX_FILE_SIZE"`
	MaxBackup int    `env:"LOG_MAX_BACKUP_NUMBER"`
	MaxAge    int    `env:"LOG_MAX_AGE"`
	Compress  bool   `env:"LOG_COMPRESS"`
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	var cnf Config
	if err := env.Parse(&cnf); err != nil {
		return nil, err
	}

	return &cnf, nil
}

func (cnf *Config) ConfigLogger() io.WriteCloser {
	return &lumberjack.Logger{
		Filename:   cnf.Logger.FilePath,
		MaxSize:    cnf.Logger.MaxSize,
		MaxBackups: cnf.Logger.MaxBackup,
		MaxAge:     cnf.Logger.MaxAge,
		Compress:   cnf.Logger.Compress,
	}

}
