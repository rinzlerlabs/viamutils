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

var components = make([]Component, 0)

func AddModularResource(api resource.API, model resource.Model) error {
	components = append(components, Component{
		API:   api,
		Model: model,
	})
	return nil
}

func RunModule(ctx context.Context, args []string, logger logging.Logger) error {
	module, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		logger.Errorf("Failed to start module: %v", err)
		return err
	}
	for _, component := range components {
		err = module.AddModelFromRegistry(ctx, component.API, component.Model)
		if err != nil {
			logger.Errorf("Failed to add model to module: %v", err)
			return err
		}
	}
	err = module.Start(ctx)
	if err != nil {
		logger.Errorf("Failed to start module: %v", err)
		return err
	}
	defer module.Close(ctx)
	<-ctx.Done()
	return nil
}
