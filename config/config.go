package config

import (
  "io/ioutil"
  "log"
  "strconv"

  "github.com/ghodss/yaml"
  "os"
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

var Config *AppConfig

func init() {
  var defaultFilePath = "src/go-get-it/local.yaml"

  goPath := os.Getenv("GOPATH")
  ReadConfig([]string{goPath + "/" + defaultFilePath})
}

/**
 * ReadConfig reads the configuration file.
 */
func ReadConfig(specifiedPaths []string)(*AppConfig){
  return readConfigWithPath(specifiedPaths)
}

func readConfigWithPath(specifiedPaths []string)(*AppConfig) {
  configFileSearcher := ConfigFileSearcher{specifiedPaths: specifiedPaths}

  filePath, err := configFileSearcher.SearchFile()
  if err != nil {
    log.Panic("File path is not valid.", err)
  }

  data, err := ioutil.ReadFile(filePath)
  if err != nil {
    log.Panic("It was not possible to read the file.", err)
  }

  unmarshallYamlData(data)
  log.Println("############################")
  log.Println("Server Config")
  log.Println("Port: " + Config.Port)
  log.Println("############################")
  log.Println("RedisClient Config")
  log.Println("host: " + Config.RedisClient.Host)
  log.Println("port: " + Config.RedisClient.Port)
  log.Println("maxIdle: " + strconv.Itoa(Config.RedisClient.MaxIdleConnections))
  log.Println("idleTimeout: " + strconv.Itoa(Config.RedisClient.IdleTimeout))
  log.Println("connectionType: " + Config.RedisClient.ConnectionType)
  log.Println("secondsToExpire: " + strconv.Itoa(Config.RedisClient.SecondsToExpire))
  log.Println("############################")
  return Config

}

func unmarshallYamlData(data []byte) {
  err := yaml.Unmarshal(data, &Config)
  if err != nil {
    log.Fatal("Error: ", err)
  }
}
