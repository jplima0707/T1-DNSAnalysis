package models

import "time"

type BenchmarkResult struct {
	Server  string
	Results []DNSResponse
	Avg     time.Duration
	Min     time.Duration
	Max     time.Duration
	Loss    float64
}
