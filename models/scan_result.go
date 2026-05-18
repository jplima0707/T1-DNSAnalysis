package models

type ScanResult struct {
	Responses []DNSResponse

	Blocked bool

	Consensus ConsensusResponse
}
