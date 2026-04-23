package main

import (
	"time"
)

// message represent a singe message
type message struct {
	Name    string
	Message string
	When    time.Time
}
