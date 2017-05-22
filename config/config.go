package config

import (
  //Go libs
  "io/ioutil"
  "log"
  "strconv"

  //External libs
  "github.com/ghodss/yaml"
)

type AppConfig struct {
  Port string
  RedisClient RedisClientConfig
}

type RedisClientConfig struct {
  Host string
  Port string
  MaxIdleConnections int
  IdleTimeout int
  ConnectionType string
  SecondsToExpire int
}

func init() {
  ReadConfig()
}

var AppConfiguration *AppConfig

func ReadConfig()(*AppConfig, error) {
  var err error
  var data []byte

  if AppConfiguration == nil {
    data, err = ioutil.ReadFile("/home/ekashivagui/workspace/src/go-get-it/config/local.yaml")
    if err != nil {
      panic(err)
    }


    err = yaml.Unmarshal(data, &AppConfiguration)
    log.Println("############################")
    log.Println("Server Config")
    log.Println("Port: " + AppConfiguration.Port)
    log.Println("############################")
    log.Println("RedisClient Config")
    log.Println("host: " + AppConfiguration.RedisClient.Host)
    log.Println("port: " + AppConfiguration.RedisClient.Port)
    log.Println("maxIdle: " + strconv.Itoa(AppConfiguration.RedisClient.MaxIdleConnections))
    log.Println("idleTimeout: " +  strconv.Itoa(AppConfiguration.RedisClient.IdleTimeout))
    log.Println("connectionType: " + AppConfiguration.RedisClient.ConnectionType)
    log.Println("secondsToExpire: " + strconv.Itoa(AppConfiguration.RedisClient.SecondsToExpire))
    log.Println("############################")
  }
  return AppConfiguration, err

}
