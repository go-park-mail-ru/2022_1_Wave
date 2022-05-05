package InitDb

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"os"
)

func InitDatabase(env string) (*sqlx.DB, error) {
	dsn := os.Getenv(env)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping() // вот тут будет первое подключение к базе
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	sqlxDb := sqlx.NewDb(db, "pgx")
	err = sqlxDb.Ping()
	if err != nil {
		return nil, err
	}

	return sqlxDb, nil
}
