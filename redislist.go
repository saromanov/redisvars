package redisvars

import
(
	"github.com/fzzy/radix/redis"
	"fmt"
)

type RedisList struct{
	//Note: Now is key - string and value - string
	inner map[string][]string
	client *redis.Client
}

func NewRedisList(client *redis.Client)*RedisList {
	value := new(RedisList)
	value.client = client
	value.inner = make(map[string][]string)
	return value
}

//Set provides setting new key value to local storage
func (dict*RedisList) SetList(key string, value []string) {
	dict.inner[key] = value
}

//Get provides getting element from redis. If element not found, returns error
func (dict*RedisList) GetList(key string)[]string {
	result, err := dict.client.Cmd("lrange", formatTitle(key), 0, -1).List()
	if err != nil {
		panic(err)
	}
	return result

}

//Commit changes
func (rv* RedisList) CommitList() {
	for key, value := range rv.inner {
		rv.client.Cmd("del", formatTitle(key))
		rv.client.Cmd("lpush", formatTitle(key), value)
	}
}

func formatTitle(name string) string{
	return fmt.Sprintf("redisvars_%s", name)
}