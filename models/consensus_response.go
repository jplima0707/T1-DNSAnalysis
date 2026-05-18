package models

type ConsensusResponse struct {
	Consensus string
	Outliers  []Outlier
}

type Outlier struct {
	Server string
	IPs    []string
}
