package models

import (
	"time"
)

type DNSResponse struct {
	Server       string
	IPs          []string
	RCode        int
	ResponseTime time.Duration
	Error        error
}
