package services

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	DB     *gorm.DB
	Redis  *redis.Client
	Logger *zap.Logger
}

func NewService() *Service {

	logger, errLogger := zap.NewProduction()
	if errLogger != nil {
		panic("Failed to handle logging")
	}

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", zap.Error(err))
	}

	defer logger.Sync() // flushes buffer, if any

	// dsn := os.Getenv("DB_DSN")
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	logger.Panic("Failed to open database : ", zap.Error(err))
	// }

	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	cache := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	})

	return &Service{nil, cache, logger}
}
