package model

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"

	"sorame/common"
)

type LinkRepository struct {
	RedisClient *redis.Client
}

func NewLinkRepository(rc *redis.Client) *LinkRepository {
	return &LinkRepository{
		RedisClient: rc,
	}
}

// InsertLink stores a new link in Redis with a 30-day expiration and returns the key
func (repo *LinkRepository) InsertLink(link *Link) (string, error) {
	ctx := context.Background()

	// Convert link to JSON
	linkJSON, err := json.Marshal(link)
	if err != nil {
		log.Println("⚠️ Error marshaling link to JSON:", err)
		return "", err
	}

	// Generate a unique ID for the link
	linkUid := common.GenerateLinkUid()

	// Store in Redis with key pattern: link:{linkUid} and 30 day expiration
	key := "link:" + linkUid
	err = repo.RedisClient.Set(ctx, key, linkJSON, 30*24*time.Hour).Err()
	if err != nil {
		log.Println("⚠️ Error storing link in Redis:", err)
		return "", err
	}

	return linkUid, nil
}

// GetLink retrieves a link from Redis by linkUid
func (repo *LinkRepository) GetLink(linkUid string) (*Link, error) {
	ctx := context.Background()

	// Get the link from Redis
	key := "link:" + linkUid
	linkJSON, err := repo.RedisClient.Get(ctx, key).Result()
	if err != nil {
		log.Println("⚠️ Error getting link from Redis:", err)
		return nil, err
	}

	var link Link
	err = json.Unmarshal([]byte(linkJSON), &link)
	if err != nil {
		log.Println("⚠️ Error unmarshaling link:", err)
		return nil, err
	}

	return &link, nil
}
