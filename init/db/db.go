package InitDb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func InitPostgres(env string) (*sqlx.DB, error) {
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

func InitMongo(urlEnv string, databaseName string, collectionName string, ctx context.Context) (*mongo.Collection, error) {
	credential := options.Credential{
		AuthMechanism:           "",
		AuthMechanismProperties: nil,
		AuthSource:              os.Getenv("MONGO_INITDB_DATABASE"),
		Username:                os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		Password:                os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
		PasswordSet:             false,
	}

	fmt.Println(credential)

	mongoUrl := os.Getenv(urlEnv)
	mongoType := os.Getenv("dbType")
	if mongoUrl == "" || mongoType == "" {
		return nil, errors.New("error to init db:" + mongoType)
	}

	clientOptions := options.Client().ApplyURI(mongoUrl).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, errors.New("error to make client to db:" + mongoType + "url:" + mongoUrl + "err:" + err.Error())
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, errors.New("error to ping to db:" + mongoType + "url:" + mongoUrl + "err:" + err.Error())
	}

	collection := client.Database(databaseName).Collection(collectionName)
	return collection, nil
}
