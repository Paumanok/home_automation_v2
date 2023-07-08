package database

import (
	"strings"
	"fmt"
	"time"
	"net"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)



type Config struct {
	Host string 	 `yaml:"Host"`
	Port string 	 `yaml:"Port"`
	Driver string 	 `yaml:"Driver"`

	SSLMode string 	 `yaml:"SSLMode"`
	
	StoreName string `yaml:"StoreName"`
	Username  string `yaml:"Username"`
	Password  string `yaml:"Password"`

	ConnPoolSize uint          `yaml:"ConnPoolSize"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
	IdleTimeout  time.Duration `yaml:"IdleTimeout"`
	DialTimeout  time.Duration `yaml:"DialTimeout"`
}


//connection url 
func (cfg *Config) connURL() string {
	sslMode := strings.TrimSpace(cfg.SSLMode)
	if sslMode == "" {
		sslMode = "disable"
	}

	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Driver, 
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.StoreName,
		sslMode,
	)
}


func NewService(cfg *Config) (*pgxpool.Pool, error) {
	poolcfg, err := pgxpool.ParseConfig(cfg.connURL())
	if err != nil {
		fmt.Println("Failed to parse config")
		return nil, err //maybe work on an error managment system
	}

	poolcfg.MaxConnLifetime = cfg.IdleTimeout
	poolcfg.MaxConns = int32(cfg.ConnPoolSize)

	dialer := &net.Dialer{KeepAlive: cfg.DialTimeout}
	dialer.Timeout = cfg.DialTimeout
	poolcfg.ConnConfig.DialFunc = dialer.DialContext

	pool, err := pgxpool.NewWithConfig(context.Background(), poolcfg)
	if err != nil {
		fmt.Println("failed to create pgx pool")
		return nil, err
	}

	return pool, nil
}
