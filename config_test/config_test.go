package config_test

import (
	"os"
	"testing"

	"go-get-it/config"
)

func TestReadConfig(t *testing.T){

	t.Log("Given the test.yaml resource.")
	{
		gopath := os.Getenv("GOPATH")

		t.Log("Should read the resource correctly.")
		{
			config.ReadConfig([]string{gopath + "/src/go-get-it/config_test/test.yaml"})

			if config.Config.Port == "9999" {
				t.Log("\t\tFound server app port")
			} else {
				t.Fatal("\t\tShould have found server app port")
			}

			if config.Config.RedisClient.Host == "storageHost" {
				t.Log("\t\tFound storage host")
			} else {
				t.Fatal("\t\tShould have found storage host")
			}

			if config.Config.RedisClient.Port == "7777" {
				t.Log("\t\tFound storage port")
			} else {
				t.Log("\t\tShould have found storage port")
			}

			if config.Config.RedisClient.MaxIdleConnections == 3 {
				t.Log("\t\tFound storage max idle connections")
			} else {
				t.Log("\t\tShould have found storage max idle connections")
			}

			if config.Config.RedisClient.IdleTimeout == 240 {
				t.Log("\t\tFound storage idle timeout")
			} else {
				t.Log("\t\tShould have found storage idle timeout")
			}

			if config.Config.RedisClient.ConnectionType == "tcp" {
				t.Log("\t\tFound storage connection type")
			} else {
				t.Log("\t\tShould have found storage connection type")
			}

			if config.Config.RedisClient.SecondsToExpire == 30 {
				t.Log("\t\tFound storage seconds to expire")
			} else {
				t.Log("\t\tShould have found storage seconds to expire")
			}
		}
	}
}
