package config

import (
	"errors"
	"go-get-it/infrastructure/logger"
	"os"
	"strings"
)

type ConfigFileSearcher struct {
	specifiedPaths []string
}

/**
 * SearchFile searches for a the first valid path specified
 */
func (fileSearcher *ConfigFileSearcher) SearchFile() (string, error) {
	for _, specifiedPath := range fileSearcher.specifiedPaths {
		if len(validatePathName(specifiedPath)) > 0 {
			if _, err := os.Stat(specifiedPath); !os.IsNotExist(err) {
				return specifiedPath, nil
			}
		}
	}
	return "Not found a valid path", errors.New("Not found a valid path")
}

func validatePathName(specifiedPath string) string {
	var folderPath string

	if len(strings.TrimSpace(specifiedPath)) > 0 {
		folderPath = specifiedPath
	} else {
		var err error
		folderPath, err = os.Getwd()
		if err != nil {
			logger.EmptyLogger().Fatal(err)
		}
	}
	return folderPath
}
