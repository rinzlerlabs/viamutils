package config

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

const (
	CloudConfigPath           = "/root/.viam/cached_cloud_config_"
	CloudConfigKey            = "cloud"
	CloudMachineIdKey         = "machine_id"
	CloudMachinePartIdKey     = "id"
	AuthKey                   = "auth"
	AuthHandlersKey           = "handlers"
	AuthHandlersConfigKey     = "config"
	AuthHandlersConfigKeysKey = "keys"
	AuthHandlersTypeKey       = "type"
)

var (
	ErrMissingCloudField              = errors.New("missing field cloud")
	ErrMissingFieldMachineId          = errors.New("missing field machine_id")
	ErrMissingFieldMachinePartId      = errors.New("missing field id")
	ErrMissingFieldAuth               = errors.New("missing field auth")
	ErrMissingFieldAuthHandlers       = errors.New("missing field handlers")
	ErrMissingFieldAuthHandlersConfig = errors.New("missing field config")
	ErrEmptyFieldAuthHandlers         = errors.New("empty field auth handlers")
	ErrAuthHandlerNotFound            = errors.New("auth handler not found")
)

func GetMachineId() (string, error) {
	config, err := GetMachineConfig()
	if err != nil {
		return "", err
	}
	if config[CloudConfigKey] == nil {
		return "", ErrMissingCloudField
	}
	cloudConfig := config[CloudConfigKey].(map[string]interface{})
	if cloudConfig[CloudMachineIdKey] == nil {
		return "", ErrMissingFieldMachineId
	}

	return cloudConfig[CloudMachineIdKey].(string), nil
}

func GetMachinePartId() (string, error) {
	config, err := GetMachineConfig()
	if err != nil {
		return "", err
	}
	if config[CloudConfigKey] == nil {
		return "", ErrMissingCloudField
	}
	cloudConfig := config[CloudConfigKey].(map[string]interface{})
	if cloudConfig[CloudMachinePartIdKey] == nil {
		return "", ErrMissingFieldMachinePartId
	}

	return cloudConfig[CloudMachinePartIdKey].(string), nil
}

func GetMachineConfig() (map[string]interface{}, error) {
	filePath, err := GetMachineConfigPath()
	if err != nil {
		return nil, err
	}
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var config map[string]interface{}
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}
	return config, nil
}

func GetMachineConfigPath() (string, error) {
	filePath := os.Getenv("VIAM_CONFIG_FILE")
	if _, err := os.Stat(filePath); os.IsNotExist(err) || filePath == "" {
		filePath = "/root/.viam/"
		files, err := os.ReadDir(filePath)
		if err != nil {
			// handle the error
			return "", err
		}
		matchingFiles := []os.DirEntry{}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if name := file.Name(); strings.HasPrefix(name, "cached_cloud_config_") && strings.HasSuffix(name, ".json") {
				matchingFiles = append(matchingFiles, file)
			}
		}

		if len(matchingFiles) != 1 {
			return "", os.ErrNotExist
		}
		return filePath + matchingFiles[0].Name(), nil
	}
	return filePath, nil
}

func GetCredentialsFromConfig() (string, string, error) {
	config, err := GetMachineConfig()
	if err != nil {
		return "", "", err
	}
	if config[AuthKey] == nil {
		return "", "", ErrMissingCloudField
	}
	authConfig := config[AuthKey].(map[string]interface{})
	if authConfig[AuthHandlersKey] == nil {
		return "", "", ErrMissingFieldAuthHandlers
	}
	handlers := authConfig[AuthHandlersKey].([]interface{})
	if len(handlers) == 0 {
		return "", "", ErrMissingFieldAuthHandlers
	}

	var handler map[string]interface{}
	for _, h := range handlers {
		handler = h.(map[string]interface{})
		if val, ok := handler[AuthHandlersTypeKey].(string); ok && val == "api-key" {
			break
		}
	}

	if handler == nil {
		return "", "", ErrAuthHandlerNotFound
	}
	handlerConfig, ok := handler[AuthHandlersConfigKey].(map[string]interface{})
	if !ok {
		return "", "", ErrMissingFieldAuthHandlersConfig
	}
	handlerConfigKeys, ok := handlerConfig[AuthHandlersConfigKeysKey].([]interface{})
	if !ok {
		return "", "", ErrMissingFieldAuthHandlersConfig
	}
	if len(handlerConfigKeys) == 0 {
		return "", "", ErrEmptyFieldAuthHandlers
	}
	keyName, ok := handlerConfigKeys[0].(string)
	if !ok {
		return "", "", ErrMissingFieldAuthHandlersConfig
	}
	key, ok := handlerConfig[keyName].(string)
	if !ok {
		return "", "", ErrMissingFieldAuthHandlersConfig
	}

	return keyName, key, nil
}
