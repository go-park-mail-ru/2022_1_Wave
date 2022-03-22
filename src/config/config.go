package config

import (
	"fmt"
	toml "github.com/pelletier/go-toml"
	"os"
)

type Config struct {
	Domain string `toml:"domain"` // Доменное имя

	SessionIDKey string `toml:"session_id"`  // как будет называться ключ сессии в cookie
	SessionKey   string `toml:"session_key"` // Криптографический ключ для сессий
	CSRFKey      string `toml:"csrf_key"`    // Криптографический ключ для CSRF защиты

	SessionExpires string `toml:"session_expires"` // время жизни сессии

	DBConnectionString string `toml:"db_connection_string"` // Строка для соединения с БД PostgreSQL

	RedisAddress string `toml:"redis_address"`
}

var C Config

// Загрузить конфигурацию из TOML файла configFilename.
func LoadConfig(configFilename string) error {
	f, err := os.Open(configFilename)
	if err != nil {
		return err
	}
	decoder := toml.NewDecoder(f).Strict(true)
	if decoder == nil {
		return fmt.Errorf("couldn't create decoder")
	}
	if err := decoder.Decode(&C); err != nil {
		return err

	}
	if C.SessionKey == "" {
		return fmt.Errorf("session_key is empty")
	}
	return nil
}
