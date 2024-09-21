package device

import "fmt"

// Credentials represents the username and password for the device.
type Credentials struct {
    Username string
    Password string
}

// Device represents a network device.
type Device struct {
    ID         string      // Unique identifier for the device (e.g., a serial number or inventory ID).
    IPAddr     string      // The IP address of the device.
    Status     string      // Status of the device (e.g., "active", "inactive").
    Role       string      // The role of the device in the network (e.g., "Router", "Switch").
    Location   string      // Physical or logical location of the device (e.g., "Data Center A").
    DeviceTag  string      // A tag to categorize or identify the device (e.g., "Core-Router").
    OSType     string      // The operating system type of the device (e.g., "Linux", "Cisco IOS").
    Credentials Credentials // Stores the login credentials (username and password).
}

// NewDevice creates and returns a new Device instance.
// It initializes the device with the given parameters such as ID, IP address, status, role, location, device tag, and OS type.
func NewDevice(id, ipAddr, status, role, location, deviceTag, osType string) *Device {
    return &Device{
        ID:        id,
        IPAddr:    ipAddr,
        Status:    status,
        Role:      role,
        Location:  location,
        DeviceTag: deviceTag,
        OSType:    osType,
    }
}

// UpdateStatus updates the current status of the network device (e.g., from "inactive" to "active").
func (d *Device) UpdateStatus(status string) {
    d.Status = status
    fmt.Printf("Device %s status updated to %s\n", d.ID, status)
}

// SetRole updates the role of the device (e.g., from "Router" to "Switch").
func (d *Device) SetRole(role string) {
    d.Role = role
    fmt.Printf("Device %s role updated to %s\n", d.ID, role)
}

// SetCredentials sets the username and password for the device.
func (d *Device) SetCredentials(username, password string) {
    d.Credentials = Credentials{
        Username: username,
        Password: password,
    }
    fmt.Printf("Credentials set for device %s (Username: %s)\n", d.ID, username)
}

// SetOSType sets the OS type (Operating System type) for the device.
func (d *Device) SetOSType(osType string) {
    d.OSType = osType
    fmt.Printf("OS Type for device %s updated to %s\n", d.ID, osType)
}

// GetOSType returns the OS type (Operating System type) of the device.
func (d *Device) GetOSType() string {
    return d.OSType
}

// DisplayDeviceDetails prints all the details of the network device.
func (d *Device) DisplayDeviceDetails() {
    fmt.Printf("Device ID: %s\n", d.ID)
    fmt.Printf("IP Address: %s\n", d.IPAddr)
    fmt.Printf("Status: %s\n", d.Status)
    fmt.Printf("Role: %s\n", d.Role)
    fmt.Printf("Location: %s\n", d.Location)
    fmt.Printf("Device Tag: %s\n", d.DeviceTag)
    fmt.Printf("OS Type: %s\n", d.OSType)
    fmt.Printf("Credentials: %s/%s\n", d.Credentials.Username, d.Credentials.Password)
}
