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
	TimeStampUTC        time.Time           `json:"timestamp_utc,omitempty"`
	ServiceName         string              `json:"service_name,omitempty"`
	ServiceProvider     string              `json:"service_provider,omitempty"`
	ServiceVersion      string              `json:"service_version,omitempty"`
	ServiceStatus       ServiceStatus       `json:"service_status,omitempty"`
	ServiceStartTimeUTC time.Time           `json:"service_starttime_utc,omitempty"`
	Uptime              float64             `json:"uptime,omitempty"`
	InboundInterfaces   []InboundInterface  `json:"inbound_interfaces,omitempty"`
	OutboundInterfaces  []OutboundInterface `json:"outbound_interfaces,omitempty"`
}

// InboundInterface is inbound network inferfaces
type InboundInterface struct {
	ApplicationName  string           `json:"application_name,omitempty"`
	ConnectionStatus ConnectionStatus `json:"connection_status,omitempty"`
	TimeStampUTC     time.Time        `json:"timestamp_utc,omitempty"`
	Hostname         string           `json:"hostname,omitempty"`
	Address          string           `json:"address,omitempty"`
	OS               string           `json:"os,omitempty"`
}

// OutboundInterface is outbound network interfaces
type OutboundInterface struct {
	ApplicationName  string           `json:"application_name,omitempty"`
	TimeStampUTC     time.Time        `json:"timestamp_utc,omitempty"`
	URLs             []string         `json:"urls,omitempty"`
	ConnectionStatus ConnectionStatus `json:"connection_status,omitempty"`
}
