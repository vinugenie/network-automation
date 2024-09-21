package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/vinugenie/network-automation/device"
)

// LoadConfig loads the configuration for a device from a JSON file.
func LoadConfig(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config map[string]string
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	fmt.Println("Configuration loaded from file:", filename)
	return config, nil
}

// SaveConfig saves the device configuration to a JSON file.
func SaveConfig(filename string, config map[string]string) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Configuration saved to file:", filename)
	return nil
}

// GenerateDeviceFeatureSet generates a map of configuration for the device based on its role and OS type. This function will be further used by another function which actually generates the config based on the parameters found in the yaml file

func GenerateDeviceFeatureSet(d *device.Device) map[string]string {
	config := make(map[string]string)

	// Basic device configuration
	config["device_id"] = d.ID
	config["ip_address"] = d.IPAddr
	config["location"] = d.Location
	config["os_type"] = d.OSType

	// Role-specific and OS-specific configuration
	switch d.Role {
	case "Router":
		switch d.OSType {
		case "IOS-XE":
			config["config_file"] = "rtr_config_ios-xe.yaml"
		case "IOS-XR":
			config["config_file"] = "rtr_config_ios-xr.yaml"
		default:
			config["config_file"] = "rtr_config.yaml"
		}

	case "Switch":
		switch d.OSType {
		case "IOS-XE":
			config["config_file"] = "sw_config_ios-xe.yaml"
		case "NX-OS":
			config["config_file"] = "sw_config_nx-os.yaml"
		default:
			config["config_file"] = "sw_config.yaml"
		}

	default:

		config["config_file"] = "linux_config.yaml"
	}

	fmt.Printf("Config file set for device %s (Role: %s, OS: %s)\n", d.ID, d.Role, d.OSType)
	return config
}

// ValidateConfig checks if the essential configurations are set for the device.
func ValidateConfig(config map[string]string) bool {
	requiredKeys := []string{"device_id", "ip_address", "os_type"}
	for _, key := range requiredKeys {
		if _, ok := config[key]; !ok {
			fmt.Printf("Missing required configuration: %s\n", key)
			return false
		}
	}
	fmt.Println("Configuration is valid.")
	return true
}

// ClearConfig resets all configurations for the device.
func ClearConfig(d *device.Device) map[string]string {
	fmt.Printf("Clearing configuration for device %s\n", d.ID)
	return make(map[string]string) // Return an empty map
}
