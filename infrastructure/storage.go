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
	Config = config.AppConfiguration
	Pool = redis.Pool{
		MaxIdle:     Config.RedisClient.MaxIdleConnections,
		IdleTimeout: time.Duration(Config.RedisClient.IdleTimeout) * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial(Config.RedisClient.ConnectionType, Config.RedisClient.Host + ":" + Config.RedisClient.Port) },
	}
}


func Save(key string, value string){
	log.Printf("Saving Value %s - %s ", key, value )

	c := Pool.Get()
	defer c.Close()

	secondsToExpire := Config.RedisClient.SecondsToExpire
	var _, err = c.Do("SETEX", key, secondsToExpire,  value)

	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func Retrieve(search string) map[string] string {
	log.Println("Retrieving Values")

	c := Pool.Get()
	defer c.Close()

	found,_:=redis.Strings(c.Do("KEYS", search))
	var returnMap = make(map[string] string)
	for _, key := range found {
		binaryValue,err:= c.Do("get", key)
		value,_:=redis.String(binaryValue, err)
		returnMap[key] = value
	}

	return returnMap
}

func delete(){

}