package config

import (
	"fmt"
	toml "github.com/pelletier/go-toml"
	"os"
)

type Config struct {
	Domain string `toml:"domain"` // Доменное имя

	SessionKey string `toml:"session_key"` // Криптографический ключ для сессий
	CSRFKey    string `toml:"csrf_key"`    // Криптографический ключ для CSRF защиты

	TokenMaxAge int64 `toml:"token_max_age"` // время жизни токена

	DBConnectionString string `toml:"db_connection_string"` // Строка для соединения с БД PostgreSQL

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
