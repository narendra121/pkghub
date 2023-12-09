package redispkg

import "github.com/redis/go-redis/v9"

type Redis struct {
	host     string
	useName  string
	password string
	protocol int
	rClient  *redis.Client
}

type RedisBuilder struct {
	redis Redis
}

func NewRedisBuilder() *RedisBuilder {
	return &RedisBuilder{redis: Redis{}}
}

func (rb *RedisBuilder) SetHost(host string) *RedisBuilder {
	rb.redis.host = host
	return rb
}

func (rb *RedisBuilder) SetUserName(userName string) *RedisBuilder {
	rb.redis.useName = userName
	return rb
}

func (rb *RedisBuilder) SetPassword(password string) *RedisBuilder {
	rb.redis.password = password
	return rb
}

func (rb *RedisBuilder) SetProtocol(version int) *RedisBuilder {
	rb.redis.protocol = version
	return rb
}

func (rb *RedisBuilder) Build() Redis {
	return rb.redis
}
