package redis_test

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/bxcodec/faker"
	"github.com/go-park-mail-ru/2022_1_Wave/websocket-server/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/websocket-server/repository/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	miniRedis := miniredis.RunT(t)

	userSyncElemsRepo := redis.NewUserSyncElemsRepo(miniRedis.Addr())
	userPlayer := domain.UserPlayerState{}
	err := faker.FakeData(&userPlayer)
	assert.NoError(t, err)

	err = userSyncElemsRepo.CreateUserPlayerState(2, &userPlayer)
	assert.NoError(t, err)

	userPlayerGet, err := userSyncElemsRepo.GetUserPlayerState(2)
	assert.NoError(t, err)
	assert.Equal(t, *userPlayerGet, userPlayer)

	err = userSyncElemsRepo.DeleteUserPlayerState(2)
	assert.NoError(t, err)

	userPlayerGet, err = userSyncElemsRepo.GetUserPlayerState(2)
	assert.Error(t, err)
	assert.Nil(t, userPlayerGet)
}
