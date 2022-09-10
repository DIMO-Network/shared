//go:generate mockgen -source redis_cache_api_service.go -destination mocks/redis_cache_api_service_mock.go -package mocks

package redis

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// CacheService combines methods of redis client and redis clustered client to have one impl that works for both, reduced to only what we use
type CacheService interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	FlushAll(ctx context.Context) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Close() error
}

// NewRedisCacheService establishes connection to Redis and creates client. db is the 0-16 db instance to use from redis.
// tls is whether redis server uses tls - false if eg. running in docker but may be true if in production. url is the redis server url.
func NewRedisCacheService(clustered bool, redisSettings Settings) CacheService {
	var tlsConfig *tls.Config
	if redisSettings.TLS {
		tlsConfig = new(tls.Config)
	}

	var r CacheService
	// handle redis cluster in prod
	if clustered {
		cc := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:     []string{redisSettings.URL},
			Password:  redisSettings.Password,
			TLSConfig: tlsConfig,
		})
		// cluster client does not have setting for DB, so must set it manually.
		cc.Do(context.Background(), fmt.Sprintf("SELECT %d", redisSettings.DB))
		r = cc
	} else {
		c := redis.NewClient(&redis.Options{
			Addr:      redisSettings.URL,
			Password:  redisSettings.Password,
			TLSConfig: tlsConfig,
			DB:        redisSettings.DB,
		})
		r = c
	}

	return r
}
