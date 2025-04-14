//go:generate mockgen -source redis_cache_api_service.go -destination mocks/redis_cache_api_service_mock.go -package mocks

package redis

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/go-redis/redis/v8"
)

// CacheService combines methods of redis client and redis clustered client to have one impl that works for both, reduced to only what we use
type CacheService interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd
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
		// cluster client does not have setting for DB b/c the SELECT command cannot be used, since Redis Cluster only supports database zero.
		r = cc
	} else {
		c := redis.NewClient(&redis.Options{
			Addr:      redisSettings.URL,
			Password:  redisSettings.Password,
			TLSConfig: tlsConfig,
		})
		r = c
	}

	return &cacheService{redisClient: r, prefix: redisSettings.KeyPrefix}
}

// cacheService is a thing wrapper over the native client to prefix the keys b/c we use the same redis server cluster for
// multiple different services due to clusters not supporting DB number.
type cacheService struct {
	redisClient CacheService
	prefix      string
}

func (cs *cacheService) Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	return cs.redisClient.Set(ctx, cs.keyBuilder(key), value, expiration)
}

func (cs *cacheService) Get(ctx context.Context, key string) *redis.StringCmd {
	return cs.redisClient.Get(ctx, cs.keyBuilder(key))
}

func (cs *cacheService) FlushAll(ctx context.Context) *redis.StatusCmd {
	return cs.redisClient.FlushAll(ctx)
}

func (cs *cacheService) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	newKeys := make([]string, len(keys))
	for i, key := range keys {
		newKeys[i] = cs.keyBuilder(key)
	}

	return cs.redisClient.Del(ctx, newKeys...)
}

func (cs *cacheService) Close() error {
	return cs.redisClient.Close()
}

// keyBuilder builds key prefixed by setting prefix. if empty just returns key
func (cs *cacheService) keyBuilder(key string) string {
	if len(cs.prefix) > 0 {
		return cs.prefix + "_" + key
	}
	return key
}
