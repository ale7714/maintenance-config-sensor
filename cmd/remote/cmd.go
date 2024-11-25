package main

import (
	"context"
	"os"

	"go.viam.com/rdk/module"

	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/config"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
	robotimpl "go.viam.com/rdk/robot/impl"
	"go.viam.com/rdk/robot/web"
	_ "go.viam.com/rdk/services/sensors/builtin"
	rdkutils "go.viam.com/rdk/utils"
	"go.viam.com/utils"

	maintenanceConfigSensor "github.com/ale7714/maintenance-config-sensor"
)

func main() {
	// NewLoggerFromArgs will create a logging.Logger at "DebugLevel" if
	// "--log-level=debug" is an argument in os.Args and at "InfoLevel" otherwise.
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("testsensor"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) (err error) {

	netconfig := config.NetworkConfig{}
	netconfig.BindAddress = "0.0.0.0:8083"

	if err := netconfig.Validate(""); err != nil {
		return err
	}

	// Update the Attributes and ConvertedAttributes with the attributes your modular resource should receive
	conf := &config.Config{
		Network: netconfig,
		Components: []resource.Config{
			{
				Name:                os.Args[1],
				API:                 sensor.API,
				Model:               maintenanceConfigSensor.Model,
				Attributes:          rdkutils.AttributeMap{},
				ConvertedAttributes: &maintenanceConfigSensor.Config{},
			},
		},
	}

	myRobot, err := robotimpl.New(ctx, conf, logger)
	if err != nil {
		return err
	}

	return web.RunWebWithConfig(ctx, myRobot, conf, logger)
}
