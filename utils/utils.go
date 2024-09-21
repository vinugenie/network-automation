package utils

import (
    "fmt"
    "os"
    "time"
    "net"
    "log"
    "os/exec"
)

// Log logs messages with a timestamp
func Log(level, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s: %s\n", timestamp, level, message)
}

// InitializeLogger initializes the logger and returns a reference to it.
func InitializeLogger(logFile string) (*os.File, error) {
    file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        return nil, err
    }
    log.SetOutput(file)
    log.Println("Logging initialized")
    return file, nil
}

// LogError logs errors to the log file.
func LogError(message string, err error) {
    if err != nil {
        log.Printf("ERROR: %s - %v", message, err)
    } else {
        log.Printf("ERROR: %s", message)
    }
}

// LogInfo logs informational messages to the log file.
func LogInfo(message string) {
    log.Printf("INFO: %s", message)
}
// IsValidIP checks if a string is a valid IP address
func IsValidIP(ip string) bool {
    return net.ParseIP(ip) != nil
}

// PingHost pings an IP address to check network connectivity
func PingHost(ip string) bool {
    cmd := exec.Command("ping", "-c", "1", ip)
    err := cmd.Run()
    return err == nil
}

// GetIPFromHostname resolves an IP address from a given hostname.
func GetIPFromHostname(hostname string) (string, error) {
    ips, err := net.LookupIP(hostname)
    if err != nil {
        return "", fmt.Errorf("could not resolve hostname %s: %v", hostname, err)
    }
    if len(ips) == 0 {
        return "", fmt.Errorf("no IPs found for hostname %s", hostname)
    }
    return ips[0].String(), nil
}

// ParseCIDR parses a CIDR notation IP address and prefix length.
func ParseCIDR(cidr string) (string, *net.IPNet, error) {
    ip, network, err := net.ParseCIDR(cidr)
    if err != nil {
        return "", nil, fmt.Errorf("invalid CIDR: %v", err)
    }
    return ip.String(), network, nil
}

// FileExists checks if a file exists
func FileExists(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

// ReadFile reads the contents of a file
func ReadFile(filename string) ([]byte, error) {
    return os.ReadFile(filename)
}

// WriteFile writes data to a file
func WriteFile(filename string, data []byte) error {
    return os.WriteFile(filename, data, 0644)
}

// HandleError logs and handles errors in a standard way
func HandleError(err error) {
    if err != nil {
        fmt.Println("Error:", err)
        // Optionally, exit the program or continue based on the severity
    }
}

// RetryOperation retries an operation a specified number of times
func RetryOperation(operation func() error, retries int) error {
    var err error
    for i := 0; i < retries; i++ {
        err = operation()
        if err == nil {
            return nil
        }
        time.Sleep(2 * time.Second) // Wait before retrying
    }
    return err
}
