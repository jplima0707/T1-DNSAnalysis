package models

import (
	"time"
)

type DNSResponse struct {
	Server       string
	IPs          []string
	RCode        string
	ResponseTime time.Duration
	Error        error
}
