# Maintenance Config Sensor modular resource

Viam Module to used to demo how maintenance configs work

Example Maintenance config 

```
"maintenance": {
    "sensor_name": "rdk:component:sensor/sensor-1",
    "maintenance_allowed_key": "disable"
  }
```

Keys that this module supports

```
"enable" : Returns true allowing machine to reconfigure
"disable": Returns false disable reconfigure
"error": Returns a non-boolean reading, disable reconfigure
```
  
