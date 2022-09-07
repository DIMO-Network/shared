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
func NewRedisCacheService(enableTLS, clustered bool, url, password string, db int) CacheService {
	var tlsConfig *tls.Config
	if enableTLS {
		tlsConfig = new(tls.Config)
	}

	var r CacheService
	// handle redis cluster in prod
	if clustered {
		cc := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:     []string{url},
			Password:  password,
			TLSConfig: tlsConfig,
		})
		// cluster client does not have setting for DB, so must set it manually.
		cc.Do(context.Background(), fmt.Sprintf("SELECT %d", db))
		r = cc
	} else {
		c := redis.NewClient(&redis.Options{
			Addr:      url,
			Password:  password,
			TLSConfig: tlsConfig,
			DB:        db,
		})
		r = c
	}

	return r
}
