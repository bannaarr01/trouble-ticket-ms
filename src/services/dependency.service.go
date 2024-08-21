package services

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/go-redis/redis"
	"log"
	"trouble-ticket-ms/src/config"
)

type AppDependencies struct {
	KeycloakClient *gocloak.GoCloak
	Context        context.Context
	KeycloakCfg    config.KeyCloakConfig
	RedisClient    *redis.Client
}

func InitAppDependencies() *AppDependencies {
	// Initialize the config
	cfg := config.New()

	// Initialize the Keycloak client
	client := gocloak.NewClient(cfg.KEYCLOAK.Host)

	// Initialize the context
	ctx := context.Background()

	redisDBAddr := fmt.Sprintf("%s:%s", cfg.REDIS.Host, cfg.REDIS.Port)

	// InitRedis initializes a new Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisDBAddr,
		Password: "",
		DB:       0, // default DB
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	log.Printf("Connected to Redis on: %v", redisDBAddr)

	return &AppDependencies{
		KeycloakClient: client,
		Context:        ctx,
		KeycloakCfg:    cfg.KEYCLOAK,
		RedisClient:    redisClient,
	}
}
