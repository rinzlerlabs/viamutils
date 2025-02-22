package config

import (
	"encoding/json"
	"errors"
	"os"

	cfg "go.viam.com/rdk/config"
	rpcUtils "go.viam.com/utils/rpc"
)

const (
	viamJsonPath              = "/etc/viam.json"
	cloudConfigPath           = "/root/.viam/cached_cloud_config_"
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
	etcViamJsonPath                   = func() string { return viamJsonPath }
	viamConfigFilePath                = func() (string, error) {
		j, err := getEtcViamJson()
		if err != nil {
			return "", err
		}
		filePath := cloudConfigPath + j.Cloud.ID + ".json"
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return "", err
		}
		return filePath, nil
	}
)

type viamJson struct {
	Cloud viamJsonCloud `json:"cloud"`
}

type viamJsonCloud struct {
	AppAddress string `json:"app_address"`
	ID         string `json:"id"`
	Secret     string `json:"secret"`
}

func getEtcViamJson() (viamJson *viamJson, err error) {
	file, err := os.ReadFile(etcViamJsonPath())
	if err != nil {
		return nil, err
	}
	viamJson = nil
	if err := json.Unmarshal(file, &viamJson); err != nil {
		return nil, err
	}
	return viamJson, nil
}

func GetMachineFqdn() (fqdn string, err error) {
	config, err := GetMachineConfig()
	if err != nil {
		return "", err
	}
	return config.Cloud.FQDN, nil
}

func GetMachineId() (machineId string, err error) {
	config, err := GetMachineConfig()
	if err != nil {
		return "", err
	}
	if config.Cloud == nil {
		return "", ErrMissingCloudField
	}
	cloudConfig := config.Cloud

	return cloudConfig.ID, nil
}

func GetMachinePartId() (partId string, err error) {
	config, err := GetMachineConfig()
	if err != nil {
		return "", err
	}
	if config.Cloud == nil {
		return "", ErrMissingCloudField
	}
	cloudConfig := config.Cloud
	if cloudConfig.ID == "" {
		return "", ErrMissingFieldMachinePartId
	}

	return cloudConfig.ID, nil
}

func GetMachineConfig() (config *cfg.Config, err error) {
	filePath, err := GetMachineConfigPath()
	if err != nil {
		return nil, err
	}
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	config = nil
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}
	return config, nil
}

func GetMachineConfigPath() (path string, err error) {
	return viamConfigFilePath()
}

func GetCredentialsFromConfig() (apiKeyID string, apiKey string, err error) {
	config, err := GetMachineConfig()
	if err != nil {
		return "", "", err
	}
	authConfig := config.Auth
	if len(authConfig.Handlers) == 0 {
		return "", "", ErrMissingFieldAuthHandlers
	}
	handlers := authConfig.Handlers

	var handler *cfg.AuthHandlerConfig
	for _, h := range handlers {
		if h.Type == rpcUtils.CredentialsTypeAPIKey {
			handler = &h
			break
		}
	}

	if handler == nil {
		return "", "", ErrAuthHandlerNotFound
	}

	handlerConfig := handler.Config
	if handlerConfig == nil {
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
