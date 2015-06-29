package redisvars

import (
	"github.com/fzzy/radix/redis"
	"time"
)

type RedisVars struct {
	client *redis.Client
	Status bool
}

//New is create redis client. If server is not found, returns panic message
func New(addr string) *RedisVars {
	rv := new(RedisVars)
	clinet, err := redis.DialTimeout("tcp", addr, time.Duration(10)*time.Second)
	if err != nil {
		panic(err)
	}
	rv.client = clinet
	rv.Status = true
	return rv
}

//NewDict provides creation of NewRedisDict object
func (rv *RedisVars) NewDict() *RedisDict {
	return NewRedisDict(rv.client)
}

func (rv *RedisVars) NewList() *RedisList {
	return NewRedisList(rv.client)
}
