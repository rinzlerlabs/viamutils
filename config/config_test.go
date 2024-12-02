package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEtcViamJsonDeserialize(t *testing.T) {
	etcViamJsonPath = func() string { return "testdata/etc_viam.json" }
	viamJson, err := getEtcViamJson()
	assert.NoError(t, err)
	assert.NotNil(t, viamJson)
	assert.NotNil(t, viamJson.Cloud)
	assert.Equal(t, viamJson.Cloud.AppAddress, "https://test.viam.com:443")
	assert.Equal(t, viamJson.Cloud.ID, "abcd")
	assert.Equal(t, viamJson.Cloud.Secret, "efgh")
}

func TestViamConfigDeserialize(t *testing.T) {
	etcViamJsonPath = func() string { return "testdata/etc_viam.json" }
	viamConfigFilePath = func() (string, error) { return "testdata/cached_cloud_config_abcd.json", nil }
	viamConfig, err := GetMachineConfig()
	assert.NoError(t, err)
	assert.NotNil(t, viamConfig)
	assert.NotNil(t, viamConfig.Cloud)
	assert.Equal(t, "abcd", viamConfig.Cloud.ID)
}
