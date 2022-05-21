package redis

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2022_1_Wave/websocket-server/domain"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

type userSyncElemsRepo struct {
	pool *redis.Pool
}

var playerStateHashTableName = "player_states"

func NewUserSyncElemsRepo(redisAddr string) domain.UserSyncPlayerRepo {
	return &userSyncElemsRepo{
		pool: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", redisAddr)
			},
		},
	}
}

func (a *userSyncElemsRepo) CreateUserPlayerState(userId uint, state *domain.UserPlayerState) error {
	client := a.pool.Get()
	defer client.Close()

	tableValue, _ := json.Marshal(state)

	_, err := client.Do("HSET", playerStateHashTableName, strconv.Itoa(int(userId)), tableValue)
	if err != nil {
		return domain.ErrSetUserPlayerState
	}

	return nil
}

func (a *userSyncElemsRepo) UpdateUserPlayerState(userId uint, state *domain.UserPlayerState) error {
	return a.CreateUserPlayerState(userId, state)
}

func (a *userSyncElemsRepo) GetUserPlayerState(userId uint) (*domain.UserPlayerState, error) {
	client := a.pool.Get()
	defer client.Close()

	dataJson, err := client.Do("HGET", playerStateHashTableName, strconv.Itoa(int(userId)))
	if err != nil || dataJson == nil {
		return nil, domain.ErrGetUserPlayerState
	}

	var state domain.UserPlayerState
	err = json.Unmarshal(dataJson.([]byte), &state)
	if err != nil {
		return nil, domain.ErrUnmarshal
	}

	return &state, nil
}

func (a *userSyncElemsRepo) DeleteUserPlayerState(userId uint) error {
	client := a.pool.Get()
	defer client.Close()

	_, err := client.Do("HDEL", playerStateHashTableName, strconv.Itoa(int(userId)))
	if err != nil {
		return domain.ErrDeletePlayerState
	}

	return nil
}
