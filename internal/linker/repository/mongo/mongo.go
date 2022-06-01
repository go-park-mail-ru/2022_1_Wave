package LinkerMongo

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type LinkerRepo struct {
	Coll *mongo.Collection
}

const (
	urlField     = "url"
	hashField    = "hash"
	counterField = "counter"
	length       = 10
)

func NewLinkerMongoRepo(coll *mongo.Collection) (domain.LinkerRepo, error) {
	idxUrl := mongo.IndexModel{
		Keys:    bson.M{urlField: 1},
		Options: options.Index().SetUnique(true),
	}

	idxHash := mongo.IndexModel{
		Keys:    bson.M{hashField: 1},
		Options: options.Index().SetUnique(true),
	}

	idxCounter := mongo.IndexModel{
		Keys:    bson.M{counterField: 1},
		Options: options.Index().SetUnique(false),
	}

	idxes := []mongo.IndexModel{idxUrl, idxHash, idxCounter}

	_, err := coll.Indexes().CreateMany(
		context.TODO(),
		idxes)

	if err != nil {
		return nil, err
	}

	return &LinkerRepo{
		Coll: coll,
	}, nil
}

type LinkInfo struct {
	Url     string
	Hash    string
	Counter int64
}

func getHash() string {
	hash := utils.RandStringRunes(length)
	return hash
}

func (collection LinkerRepo) GetByField(field string, value string) (*LinkInfo, error) {
	fromMongo := LinkInfo{}
	filter := bson.M{field: value}
	if err := collection.Coll.FindOne(context.TODO(), filter).Decode(&fromMongo); err != nil {
		return nil, err
	}
	logger.GlobalLogger.Logrus.Warnln("success to get ", field, ": value: ", value)
	return &fromMongo, nil
}

func (collection LinkerRepo) GetByHash(hash string) (*LinkInfo, error) {
	return collection.GetByField(hashField, hash)
}

func (collection LinkerRepo) GetByUrl(url string) (*LinkInfo, error) {
	return collection.GetByField(urlField, url)
}

func (collection LinkerRepo) Create(url string) (string, error) {
	fromMongo, err := collection.GetByUrl(url)
	if err == nil {
		logger.GlobalLogger.Logrus.Warnln("success to get hash:", fromMongo.Hash, "for url:", fromMongo.Url)
		return fromMongo.Hash, err
	}

	hash := getHash()
	logger.GlobalLogger.Logrus.Warn("check hash:", hash)

	for collection.Coll.FindOne(context.TODO(), bson.M{hashField: hash}) == nil {
		logger.GlobalLogger.Logrus.Warn("check hash:", hash)
		hash = getHash()
	}

	data := LinkInfo{
		Url:     url,
		Hash:    hash,
		Counter: int64(0),
	}

	insertResult, err := collection.Coll.InsertOne(context.TODO(), data)
	if err != nil {
		logger.GlobalLogger.Logrus.Warn("error to insert:", err, " data:", data)
		return "", err
	}

	logger.GlobalLogger.Logrus.Info("success insert single document:", insertResult, " id:", insertResult.InsertedID)
	return data.Hash, nil
}

func (collection LinkerRepo) Get(hash string) (string, error) {
	logger.GlobalLogger.Logrus.Info("in get url by hash...")
	filter := bson.M{hashField: hash}
	update := bson.M{
		"$inc": bson.M{counterField: 1},
	}

	// Create an instance of an options and set the desired options
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	// Find one result and update it
	result := collection.Coll.FindOneAndUpdate(context.TODO(), filter, update, &opt)
	err := result.Err()
	if err != nil {
		logger.GlobalLogger.Logrus.Warnln("error to get:", err, " hash:", hash)
		return "", err
	}

	// Decode the result
	fromCollection := LinkInfo{}
	decodeErr := result.Decode(&fromCollection)
	url := fromCollection.Url
	counter := fromCollection.Counter

	logger.GlobalLogger.Logrus.Info("success get url by hash:", hash, " url:", url, " counter:", counter)
	return url, decodeErr

}

func (collection LinkerRepo) Count(hash string) (int64, error) {
	logger.GlobalLogger.Logrus.Info("in count url by hash...")

	fromMongo, err := collection.GetByHash(hash)
	if err != nil {
		return 0, err
	}

	logger.GlobalLogger.Logrus.Info("success get counter by hash:", hash, " counter:", fromMongo.Counter)
	return fromMongo.Counter, nil

}
