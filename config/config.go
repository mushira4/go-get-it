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
  Port int
  Log string
}

func ReadConfig()(AppConfig, error) {
  data, err := ioutil.ReadFile("/home/ekashivagui/workspace/src/go-get-it/config/local.yaml")
  if err != nil {
    panic(err)
  }

  var appConfig AppConfig
  err = yaml.Unmarshal(data, &appConfig)
  log.Println("############################")
  log.Println("Server Config")
  log.Println("Port: " + strconv.Itoa(appConfig.Port))
  log.Println("Log Path: " + appConfig.Log)
  log.Println("############################")

  return appConfig, err

}
