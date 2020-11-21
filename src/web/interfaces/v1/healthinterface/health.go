package healthinterface

import "time"

// ServiceStatus represents the service status
type ServiceStatus string

// ConnectionStatus represents the connection status
type ConnectionStatus string

var (
	// ServiceRunning represents service is Running
	ServiceRunning ServiceStatus = "Running"
	// ServiceDegraded represents service health is Degraded
	ServiceDegraded ServiceStatus = "Degraded"
	// ServiceStopped represents service is Stopped
	ServiceStopped ServiceStatus = "Stopped"
	// ConnectionActive represents connection is active
	ConnectionActive ConnectionStatus = "Active"
	// ConnectionDisconnected represents connection is disconnected
	ConnectionDisconnected ConnectionStatus = "Disconnected"
)

// Health represents health response
type Health struct {
	TimeStampUTC        time.Time           `json:"timeStampUTC,omitempty"`
	ServiceName         string              `json:"serviceName,omitempty"`
	ServiceProvider     string              `json:"serviceProvider,omitempty"`
	ServiceVersion      string              `json:"serviceVersion,omitempty"`
	ServiceStatus       ServiceStatus       `json:"serviceStatus,omitempty"`
	ServiceStartTimeUTC time.Time           `json:"serviceStartTimeUTC,omitempty"`
	Uptime              float64             `json:"uptime,omitempty"`
	InboundInterfaces   []InboundInterface  `json:"inboundInterfaces,omitempty"`
	OutboundInterfaces  []OutboundInterface `json:"outboundInterfaces,omitempty"`
}

// InboundInterface is inbound network inferfaces
type InboundInterface struct {
	ApplicationName  string           `json:"applicationName,omitempty"`
	ConnectionStatus ConnectionStatus `json:"connectionStatus,omitempty"`
	TimeStampUTC     time.Time        `json:"timeStampUTC,omitempty"`
	Hostname         string           `json:"hostname,omitempty"`
	Address          string           `json:"address,omitempty"`
	OS               string           `json:"os,omitempty"`
}

// OutboundInterface is outbound network interfaces
type OutboundInterface struct {
	ApplicationName  string           `json:"applicationName,omitempty"`
	TimeStampUTC     time.Time        `json:"timeStampUTC,omitempty"`
	URLs             []string         `json:"urls,omitempty"`
	ConnectionStatus ConnectionStatus `json:"connectionStatus,omitempty"`
}
