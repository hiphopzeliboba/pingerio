package model

import "time"

type Container struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Image    string    `json:"image"`
	Status   string    `json:"status"`
	Created  time.Time `json:"created"`
	PingTime time.Time `json:"ping_time"`
}

//ID        string    `json:"id"`
//Name      string    `json:"name"`
//Image     string    `json:"image"`
//Status    string    `json:"status"`
//Ports     []string  `json:"ports"`
//Reachable bool      `json:"reachable"`
//PingTime  time.Time `json:"ping_time"`
