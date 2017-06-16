package scraff

import (
	"github.com/go-redis/redis"
)

type RedisAdStore struct {
	redisClient *redis.Client
}

func NewRedisAdStore(redisURL string) AdStore {
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisURL,
		DB:   0,
	})

	return &RedisAdStore{
		redisClient: redisClient,
	}
}

func (r RedisAdStore) IsMissing(ad Ad) (b bool, err error) {
	b, err = r.redisClient.SIsMember("scraff", ad.Url).Result()
	return
}

func (r RedisAdStore) Add(ad Ad) (err error) {
	_, err = r.redisClient.SAdd("scraff", ad.Url).Result()
	return
}
