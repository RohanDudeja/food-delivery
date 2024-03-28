package common

import (
	"log"
	"strconv"
	"sync"

	"food-delivery/internal/infra/config"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type Storage struct {
	DB    *gorm.DB
	Redis *redis.UniversalClient
}

func bindStorage(app *App) error {
	var err error
	app.Storage.DB, err = initialiseDB(app.Config.DbURL())
	if err != nil {
		log.Printf("InitialiseDB error: %v", err)
		return err
	}
	app.Storage.Redis, err = initialiseRedis(app.Config)
	if err != nil {
		log.Printf("InitialiseRedis error: %v", err)
		return err
	}
	return nil
}

// InitialiseDB ...assign connection to global *gorm.DB variable DB
func initialiseDB(dbString string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dbString)
	if err != nil {
		log.Printf("DB: failed to create client: %v\n", err)
		return nil, err
	}
	if gin.IsDebugging() {
		db.LogMode(true)
	}
	return db, nil
}

func initialiseRedis(config *config.Config) (*redis.UniversalClient, error) {
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{config.Redis.Host + ":" + strconv.Itoa(config.Redis.Port)},
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
	pong, err := redisClient.Ping().Result()
	if err != nil {
		log.Printf("Redis: failed to create client: %v\n", err)
		return nil, err
	}
	log.Println(pong)
	return &redisClient, nil
}

func NewMutex() *sync.Mutex {
	return &sync.Mutex{}
}
