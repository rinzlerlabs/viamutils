package module

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

func Test_SetConfig(t *testing.T) {
	model := resource.NewModel("test", "test", "test")
	api := sensor.API
	assert.Equal(t, len(config.Components), 0)

	AddModularResource(api, model)
	assert.Equal(t, len(config.Components), 1)
	assert.Equal(t, config.Components[0].Model, model)
	assert.Equal(t, config.Components[0].API, api)
}

func Test_DisableErrorOnAdd(t *testing.T) {
	logger := logging.NewTestLogger(t)
	model := resource.NewModel("test", "test", "test")
	api := sensor.API
	SetConfig(StartupConfig{
		Components: make([]Component, 0),
		FailOnAdd:  false,
	})
	assert.Equal(t, len(config.Components), 0)

	err := AddModularResource(api, model)
	assert.Nil(t, err)

	ctx := context.Background()
	err = RunModule(ctx, []string{}, logger)
	assert.Nil(t, err)
	ctx.Done()
}

func Test_ErrorOnAdd(t *testing.T) {
	resourceError := "resource with API %s and model %s not yet registered"
	logger := logging.NewTestLogger(t)
	model := resource.NewModel("test", "test", "test")
	api := sensor.API
	SetConfig(StartupConfig{
		Components: make([]Component, 0),
		FailOnAdd:  true,
	})
	assert.Equal(t, len(config.Components), 0)

	err := AddModularResource(api, model)
	assert.Nil(t, err)

	ctx := context.Background()
	err = RunModule(ctx, []string{}, logger)
	assert.Equal(t, err.Error(), errors.Errorf(resourceError, api, model).Error())
	ctx.Done()
}
