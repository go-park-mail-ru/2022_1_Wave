package LinkerMongo

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type LinkerRepo struct {
	Coll *mongo.Collection
}

func NewLinkerMongoRepo(coll *mongo.Collection) (domain.LinkerRepo, error) {
	idxUrl := mongo.IndexModel{
		Keys:    bson.M{"url": 1},
		Options: options.Index().SetUnique(true),
	}

	idxHash := mongo.IndexModel{
		Keys:    bson.M{"hash": 1},
		Options: options.Index().SetUnique(true),
	}

	idxes := make([]mongo.IndexModel, 2)
	idxes[0] = idxUrl
	idxes[1] = idxHash

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

type LinkerPair struct {
	Url  string
	Hash string
}

func (collection LinkerRepo) Create(url string) (string, error) {
	fromMongo := LinkerPair{}
	filter := bson.M{"url": url}
	if err := collection.Coll.FindOne(context.TODO(), filter).Decode(&fromMongo); err == nil {
		logger.GlobalLogger.Logrus.Warnln("success to get hash:", fromMongo.Hash, "for url:", fromMongo.Url)
		return fromMongo.Hash, err
	}

	data := LinkerPair{
		Url:  url,
		Hash: uuid.NewString()[:7],
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
	result := LinkerPair{}
	filter := bson.M{"hash": hash}
	if err := collection.Coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		logger.GlobalLogger.Logrus.Warnln("error to get:", err, " hash:", hash)
		return "", err
	}
	url := result.Url
	logger.GlobalLogger.Logrus.Info("success get url by hash:", hash, " url:", url)
	return url, nil
}
