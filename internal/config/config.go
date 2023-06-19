package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"time"
)

type (
	Configs struct {
		Postgres Postgres
		Http     Http
	}

	Postgres struct {
		Host     string
		Port     string
		Username string
		Password string
		DbName   string
		SslMode  string
	}

	Http struct {
		Host               string
		Port               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		MaxHeaderMegabytes int
	}
)

func (c Postgres) DataSource() string {
	format := "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	return fmt.Sprintf(format, c.Host, c.Port, c.Username, c.Password, c.DbName, c.SslMode)
}

func Init(configPath string) (*Configs, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("main")

	if error := viper.ReadInConfig(); error != nil {
		return nil, error
	}

	if error := godotenv.Load(); error != nil {
		return nil, error
	}

	postgresConfig := Postgres{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("postgresPassword"),
		DbName:   viper.GetString("db.dbname"),
		SslMode:  viper.GetString("db.sslmode"),
	}

	http := Http{
		Host:               viper.GetString("http.host"),
		Port:               viper.GetString("http.port"),
		ReadTimeout:        viper.GetDuration("http.readTimeout"),
		WriteTimeout:       viper.GetDuration("http.writeTimeout"),
		MaxHeaderMegabytes: viper.GetInt("http.maxHeaderBytes"),
	}

	return &Configs{Postgres: postgresConfig, Http: http}, nil
}
