package redispkg

import "github.com/redis/go-redis/v9"

type RedisCfg struct {
	host     string
	useName  string
	password string
	protocol int
	rClient  *redis.Client
}

type RedisCfgBuilder struct {
	redisCfg RedisCfg
}

func NewRedisCfgBuilder() *RedisCfgBuilder {
	return &RedisCfgBuilder{redisCfg: RedisCfg{}}
}

func (rb *RedisCfgBuilder) SetHost(host string) *RedisCfgBuilder {
	rb.redisCfg.host = host
	return rb
}

func (rb *RedisCfgBuilder) SetUserName(userName string) *RedisCfgBuilder {
	rb.redisCfg.useName = userName
	return rb
}

func (rb *RedisCfgBuilder) SetPassword(password string) *RedisCfgBuilder {
	rb.redisCfg.password = password
	return rb
}

func (rb *RedisCfgBuilder) SetProtocol(version int) *RedisCfgBuilder {
	rb.redisCfg.protocol = version
	return rb
}

func (rb *RedisCfgBuilder) Build() RedisCfg {
	return rb.redisCfg
}
