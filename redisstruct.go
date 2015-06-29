package redisvars

import
(
	"github.com/fzzy/radix/redis"
	"encoding/json"
)

type RedisStruct struct {
	inner  map[string]string
	client *redis.Client
}

//NewRedisStruct provides create on the new object
func NewRedisStruct(client *redis.Client) *RedisStruct {
	value := new(RedisStruct)
	value.client = client
	value.inner = map[string]string{}
	return value
}

func (rs *RedisStruct) SetStruct(key string, strobj interface{}){
	res, err := json.Marshal(strobj)
    if err != nil {
        panic(err)
    }

    rs.inner[key] = string(res)
}

//Get provides getting element from redis. If element not found, returns error
func (dict *RedisStruct) GetStruct(key string) string {
	return dict.client.Cmd("hget", "redisstruct", key).String()

}

func (rv *RedisStruct) CommitStruct() {
	for key, value := range rv.inner {
		rv.client.Cmd("hset", "redisstruct", key, value)
	}
}