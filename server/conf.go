package main

type ApplicationConfiguration struct {
	ServerConfig *serverConfiguration `json:"server"`
}

type serverConfiguration struct {
	Port *string `json:"port"`
}