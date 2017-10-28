package config

import (
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {

	t.Log("Given the test.yaml resource.")
	{
		gopath := os.Getenv("GOPATH")

		t.Log("Should read the resource correctly.")
		{
			ReadConfig([]string{gopath + "/src/go-get-it/config/test.yaml"})

			if Config.Port == "9999" {
				t.Log("\t\tFound server app port")
			} else {
				t.Fatal("\t\tShould have found server app port")
			}

			if Config.RedisClient.Host == "storageHost" {
				t.Log("\t\tFound storage host")
			} else {
				t.Fatal("\t\tShould have found storage host")
			}

			if Config.RedisClient.Port == "7777" {
				t.Log("\t\tFound storage port")
			} else {
				t.Log("\t\tShould have found storage port")
			}

			if Config.RedisClient.MaxIdleConnections == 3 {
				t.Log("\t\tFound storage max idle connections")
			} else {
				t.Log("\t\tShould have found storage max idle connections")
			}

			if Config.RedisClient.IdleTimeout == 240 {
				t.Log("\t\tFound storage idle timeout")
			} else {
				t.Log("\t\tShould have found storage idle timeout")
			}

			if Config.RedisClient.ConnectionType == "tcp" {
				t.Log("\t\tFound storage connection type")
			} else {
				t.Log("\t\tShould have found storage connection type")
			}

			if Config.RedisClient.SecondsToExpire == 30 {
				t.Log("\t\tFound storage seconds to expire")
			} else {
				t.Log("\t\tShould have found storage seconds to expire")
			}
		}
	}
}
