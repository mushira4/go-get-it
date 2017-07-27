package config_test

import (
	"os"
	"io/ioutil"
	"testing"

	"go-get-it/config"
	"github.com/ghodss/yaml"
)

func TestReadConfig(t * testing.T){

	t.Log("Given the test.yaml resource.")
	{
		gopath := os.Getenv("GOPATH")
		data, _ := ioutil.ReadFile(gopath + "/src/go-get-it/config_test/test.yaml")

		t.Log("Should read the resource correctly.")
		{
			yaml.Unmarshal(data, &config.AppConfiguration)

			if config.AppConfiguration.Port == "9999" {
				t.Log("\t\tFound server app port")
			} else {
				t.Fatal("\t\tShould have found server app port")
			}

			if config.AppConfiguration.RedisClient.Host == "storageHost" {
				t.Log("\t\tFound storage host")
			} else {
				t.Fatal("\t\tShould have found storage host")
			}

			if config.AppConfiguration.RedisClient.Port == "7777" {
				t.Log("\t\tFound storage port")
			} else {
				t.Log("\t\tShould have found storage port")
			}

			if config.AppConfiguration.RedisClient.MaxIdleConnections == 3 {
				t.Log("\t\tFound storage max idle connections")
			} else {
				t.Log("\t\tShould have found storage max idle connections")
			}

			if config.AppConfiguration.RedisClient.IdleTimeout == 240 {
				t.Log("\t\tFound storage idle timeout")
			} else {
				t.Log("\t\tShould have found storage idle timeout")
			}

			if config.AppConfiguration.RedisClient.ConnectionType == "tcp" {
				t.Log("\t\tFound storage connection type")
			} else {
				t.Log("\t\tShould have found storage connection type")
			}

			if config.AppConfiguration.RedisClient.SecondsToExpire == 30 {
				t.Log("\t\tFound storage seconds to expire")
			} else {
				t.Log("\t\tShould have found storage seconds to expire")
			}
		}
	}
}
