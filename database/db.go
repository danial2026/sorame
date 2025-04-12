package database

import (
	"context"
	"log"
	"os"
	"strconv"

	"sorame/model"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() *model.LinkRepository {
	// Get Redis configuration from environment variables
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	dbStr := os.Getenv("REDIS_DB")

	// Convert DB string to integer
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		log.Fatalf("⚠️ Error converting REDIS_DB to integer: %v", err)
	}

	// Create Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})

	// Test the connection
	ctx := context.Background()
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("⚠️ Error connecting to Redis: %v", err)
	}

	log.Println("✅ Successfully connected to Redis")

	return model.NewLinkRepository(
		RedisClient,
	)
}
