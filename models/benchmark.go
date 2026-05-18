package models

import "time"

type BenchmarkResult struct {
	ServerName string
	ServerIP   string
	Results    []DNSResponse
	Avg        time.Duration
	Min        time.Duration
	Max        time.Duration
	Loss       float64
}
