package main

type ApplicationConfiguration struct {
	ServerConfig *ServerConfiguration `json:"server"`
}

type ServerConfiguration struct {
	Port *string `json:"port"`
}

type DataBaseConfiguration struct {
	Driver             string `json:"driver"`
	User               string `json:"user"`
	Password           string `json:"password"`
	Host               string `json:"host"`
	Port               int    `json:"port"`
	Database           string `json:"database"`
	MaxIdleConnections int    `json:"max_idle_connections"`
	MaxOpenConnections int    `json:"max_open_connections"`
	ApplicationName    string `json:"application_name"`
	SSLMode            string `json:"ssl_mode"`
}
