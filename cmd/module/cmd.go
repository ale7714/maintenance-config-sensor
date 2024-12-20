// Package main is a module which serves the customsensor custom model.
package main

import (
	"context"

	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"

	maintenanceConfigSensor "github.com/ale7714/maintenance-config-sensor"
)

func main() {
	// NewLoggerFromArgs will create a logging.Logger at "DebugLevel" if
	// "--log-level=debug" is an argument in os.Args and at "InfoLevel" otherwise.
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("maintenance-config-sensor"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) (err error) {
	myModule, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	// Adds the preregistered sensor component API to the module for the new model.
	// TODO: Update the name of your package customsensor
	err = myModule.AddModelFromRegistry(ctx, sensor.API, maintenanceConfigSensor.Model)
	if err != nil {
		return err
	}

	err = myModule.Start(ctx)
	defer myModule.Close(ctx)
	if err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}
