package scraff

import "github.com/go-redis/redis"
import "log"

type AdStore interface {
	//IsPresent(ad Ad)
	update(key, url string, weigth float64)
}

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

func (r RedisAdStore) update(key, url string, weigth float64) {
	_, err := r.redisClient.ZIncr(key, redis.Z{weigth, url}).Result()
	if err != nil {
		log.Println(err)
	}
}
