package redisvars

import (
	"github.com/fzzy/radix/redis"
)

type RedisDict struct {
	//Note: Now is key - string and value - string
	inner  map[string]string
	client *redis.Client
}

func NewRedisDict(client *redis.Client) *RedisDict {
	value := new(RedisDict)
	value.client = client
	value.inner = make(map[string]string)
	return value
}

//Set provides setting new key value to local storage
func (dict *RedisDict) Set(key, value string) {
	dict.inner[key] = value
}

//Get provides getting element from redis. If element not found, returns error
func (dict *RedisDict) Get(key string) string {
	return dict.client.Cmd("hget", "redisvars", key).String()

}

func (dict *RedisDict) Delete(key string) {
	dict.client.Cmd("hdel", "redisvars", key)
}

func (rv *RedisDict) Commit() {
	for key, value := range rv.inner {
		rv.client.Cmd("hset", "redisvars", key, value)
	}
}
