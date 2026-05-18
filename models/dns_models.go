package models

import "time"

type DNSRecord struct {
	Type uint16
	TTL  uint32
	Data string
}

type DNSFlags struct {
	Response           bool
	Authoritative      bool
	Truncated          bool
	RecursionDesired   bool
	RecursionAvailable bool
}

type ParsedDNSResponse struct {
	ID          uint16
	RCode       int
	Flags       DNSFlags
	QuestionCnt int
	AnswerCnt   int
	Records     []DNSRecord
}

type QueryResult struct {
	Attempt      int
	Protocol     string
	ResponseTime time.Duration
	Response     *ParsedDNSResponse
	Error        string
}

type ServerReport struct {
	ServerName string
	Provider   string
	ServerIP   string
	Protocol   string
	Attempts   int
	Successes  int
	Losses     int
	LossRate   float64
	Avg        time.Duration
	Min        time.Duration
	Max        time.Duration
	IPs        []string
	RCodes     map[int]int
	Findings   []string
	Results    []QueryResult
}
