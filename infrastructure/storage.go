package infrastructure

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"

	"go-get-it/config"
)

var Pool redis.Pool
var Config *config.AppConfig

func init(){
	Config = config.Config
	Pool = redis.Pool{
		MaxIdle:     Config.RedisClient.MaxIdleConnections,
		IdleTimeout: time.Duration(Config.RedisClient.IdleTimeout) * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial(Config.RedisClient.ConnectionType, Config.RedisClient.Host + ":" + Config.RedisClient.Port) },
	}
}


func Save(key string, value string){
	c := Pool.Get()
	defer c.Close()

	secondsToExpire := Config.RedisClient.SecondsToExpire
	var _, err = c.Do("SETEX", key, secondsToExpire,  value)

	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func Retrieve(searchQuery string) map[string] string {
	c := Pool.Get()
	defer c.Close()

	foundValues, err :=redis.Strings(c.Do("KEYS", searchQuery))
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var returnMap = make(map[string] string)
	for _, key := range foundValues {
		binaryValue,err:= c.Do("get", key)
		value,_:=redis.String(binaryValue, err)
		returnMap[key] = value
	}

	return returnMap
}