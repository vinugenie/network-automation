package main

import (
	"fmt"
    "github.com/vinugenie/network-automation/device"
    "github.com/vinugenie/network-automation/config"
	// "github.com/vinugenie/network-automation/utils"
    "github.com/vinugenie/network-automation/monitor"
)

func main() {
	// Create a new device with all necessary details
	d := device.NewDevice(
		"1",                // Device ID
		"192.168.1.10",     // IP Address
		"inactive",         // Status
		"Router",           // Role
		"Data Center A",    // Location
		"Core-Router",      // Device Tag
		"IOS-XE",           // OS Type
	)

	// Set the device credentials
	d.SetCredentials("admin", "password123")

	// Generate a configuration for the device based on its role and OS type
	deviceConfig := config.GenerateDeviceFeatureSet(d)

	// Save the generated configuration to a file
	err := config.SaveConfig("device-config.json", deviceConfig)
	if err != nil {
		fmt.Println("Error saving configuration:", err)
		return
	}

	// Load the configuration from the file to verify it's correctly saved
	loadedConfig, err := config.LoadConfig("device-config.json")
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	// Validate the loaded configuration
	if config.ValidateConfig(loadedConfig) {
		fmt.Println("Loaded configuration is valid.")
	}

	// Log device details
	d.DisplayDeviceDetails()

	// Monitor device health by pinging it
	isHealthy := monitor.CheckDeviceHealth(d.IPAddr)
	if isHealthy {
		fmt.Println("Device is healthy!")
	} else {
		fmt.Println("Device is unreachable!")
	}

	// Set alert thresholds for CPU, memory, and bandwidth usage
	monitor.SetAlertThresholds(d.IPAddr, 80.0, 75.0, 800.0)

	// Simulate monitoring CPU, memory, and traffic usage
	cpuUsage := monitor.MonitorCPUUsage(d.IPAddr)
	memoryUsage := monitor.MonitorMemoryUsage(d.IPAddr)
	traffic := monitor.MonitorNetworkTraffic(d.IPAddr)

	// Log device metrics
	monitor.LogDeviceMetrics(d.IPAddr, cpuUsage, memoryUsage, traffic["inbound"], traffic["outbound"])

	// Monitor thresholds and trigger alerts if any limits are exceeded
	monitor.MonitorThresholdsAndAlert(d.IPAddr, cpuUsage, memoryUsage, traffic["inbound"], traffic["outbound"])

	// Get device uptime and check if it has rebooted recently
	uptime := monitor.GetDeviceUptime(d.IPAddr)
	fmt.Printf("Device uptime: %s\n", uptime)

	rebooted := monitor.CheckDeviceRebootStatus(d.IPAddr)
	if rebooted {
		fmt.Println("Device has rebooted recently.")
	} else {
		fmt.Println("Device has not rebooted.")
	}

	// Clear the device configuration (simulating decommissioning)
	clearedConfig := config.ClearConfig(d)
	fmt.Println("Cleared configuration:", clearedConfig)
}
