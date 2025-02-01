package module

import (
	"context"

	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/resource"
)

type Component struct {
	API   resource.API
	Model resource.Model
}

type StartupConfig struct {
	Components []Component
	FailOnAdd  bool
}

var config = StartupConfig{
	Components: make([]Component, 0),
	FailOnAdd:  true,
}

func SetConfig(c StartupConfig) {
	if c.Components == nil {
		c.Components = make([]Component, 0)
	}
	config = c
}

func AddModularResource(api resource.API, model resource.Model) error {
	config.Components = append(config.Components, Component{
		API:   api,
		Model: model,
	})
	return nil
}

func RunModule(ctx context.Context, args []string, logger logging.Logger) error {
	logger.Infof("Starting Viam Module with args: %v", args)
	module, err := module.NewModuleFromArgs(ctx)
	if err != nil {
		logger.Errorf("Failed to start module: %v", err)
		return err
	}
	logger.Infof("Module started successfully, registering %v models", len(config.Components))
	for _, component := range config.Components {
		logger.Infof("Adding %v to module with API %v", component.Model, component.API)
		err = module.AddModelFromRegistry(ctx, component.API, component.Model)
		if err != nil {
			logger.Errorf("Failed to add model (%v) to module: %v", component.Model, err)
			if config.FailOnAdd {
				return err
			}
		}
	}
	logger.Info("Starting module")
	err = module.Start(ctx)
	if err != nil {
		logger.Errorf("Failed to start module: %v", err)
		return err
	}
	defer module.Close(ctx)
	logger.Info("Module started successfully")
	<-ctx.Done()
	logger.Info("Module stopped")
	return nil
}
