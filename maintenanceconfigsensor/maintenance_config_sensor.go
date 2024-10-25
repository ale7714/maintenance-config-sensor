package maintenanceConfigSensor

import (
	"context"
	"errors"

	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

var (
    Model = resource.NewModel("viam-labs", "maintenance-config-sensor", "maintenanceConfigSensor")
    errUnimplemented = errors.New("unimplemented")
)

func init() {
    resource.RegisterComponent(sensor.API, Model,
        resource.Registration[sensor.Sensor, *Config]{
            Constructor: newMaintenanceConfigSensor,
        },
    )
}

type Config struct {}

// Validate validates the config and returns implicit dependencies.
func (cfg *Config) Validate(path string) ([]string, error) {
    return []string{}, nil
}

// Constructor for a custom sensor that creates and returns a customSensor.
// TODO: update the customSensor struct and the initialization.
func newMaintenanceConfigSensor(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (sensor.Sensor, error) {
    // This takes the generic resource.Config passed down from the parent and converts it to the
    // model-specific (aka "native") Config structure defined above, making it easier to directly access attributes.
    conf, err := resource.NativeConfig[*Config](rawConf)
    if err != nil {
        return nil, err
    }

    // Create a cancelable context for custom sensor
    cancelCtx, cancelFunc := context.WithCancel(context.Background())

    s := &maintenanceConfigSensor{
        name:        rawConf.ResourceName(),
        logger:      logger,
        cfg:         conf,
        cancelCtx:   cancelCtx,
        cancelFunc:  cancelFunc,
    }

    // The Reconfigure() method changes the values on the customSensor based on the attributes in the component config
    if err := s.Reconfigure(ctx, deps, rawConf); err != nil {
        logger.Error("Error configuring module with ", rawConf)
        return nil, err
    }

    return s, nil
}

type maintenanceConfigSensor struct {
    name   resource.Name
    logger logging.Logger
    cfg    *Config

    cancelCtx  context.Context
    cancelFunc func()
}

func (s *maintenanceConfigSensor) Name() resource.Name {
    return s.name
}

// Reconfigures the model. Most models can be reconfigured in place without needing to rebuild. If you need to instead create a new instance of the sensor, throw a NewMustBuildError.
func (s *maintenanceConfigSensor) Reconfigure(ctx context.Context, deps resource.Dependencies, conf resource.Config) error {
    s.name = conf.ResourceName()
    
    return nil
}

func (s *maintenanceConfigSensor) Readings(ctx context.Context, extra map[string]interface{}) (map[string]interface{}, error) {
    return map[string]interface{}{"enable": true, "disable": false, "error": 1}, nil
}

// DoCommand is a place to add additional commands to extend the sensor API. This is optional.
func (s *maintenanceConfigSensor) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
    s.logger.Error("DoCommand method unimplemented")
    return nil, errUnimplemented
}

// Close closes the underlying generic.
func (s *maintenanceConfigSensor) Close(ctx context.Context) error {
    s.cancelFunc()
    return nil
}
