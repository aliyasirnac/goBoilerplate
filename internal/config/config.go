package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aliyasirnac/goBackendBoilerplate/internal/loggerx"
	"github.com/joho/godotenv"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	koanfyaml "github.com/knadh/koanf/v2"
	"github.com/pkg/errors"
)

type Config struct {
	App      App
	Database Database
	Postgres Postgres
}

type App struct {
	Port int
	Log  loggerx.Config
}

type Postgres struct {
	Dsn string
}

type Database struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	SslMode  string
}

func LoadConfig() (*Config, error) {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		return nil, errors.Wrap(err, "while loading .env file")
	}

	// YAML dosyasını yükle
	k := koanfyaml.New(".")
	parser := yaml.Parser()
	configFile := "config.yml"

	// YAML dosyasını oku ve env değişkenleriyle değiştir
	yamlContent, err := os.ReadFile(configFile)
	if err != nil {
		return nil, errors.Wrap(err, "while reading config file")
	}

	// İçeriği string olarak al ve env değişkenleriyle değiştir
	yamlString := os.ExpandEnv(string(yamlContent))

	// Değiştirilmiş içeriği tekrar yaml provider'a ver
	if err := k.Load(rawbytes.Provider([]byte(yamlString)), parser); err != nil {
		return nil, errors.Wrap(err, "while loading config file with replaced env vars")
	}

	var cfg Config
	err = k.Unmarshal("", &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "while unmarshalling config file")
	}

	cfg.Database.Port, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, errors.Wrap(err, "while parsing DB_PORT")
	}

	cfg.Postgres = *NewPostgres(cfg.Database)

	return &cfg, nil
}

func NewPostgres(d Database) *Postgres {
	return &Postgres{
		Dsn: formatDbConnStr(d),
	}
}

func formatDbConnStr(d Database) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", d.Host, d.User, d.Password, d.DBName, d.Port, d.SslMode)
}
