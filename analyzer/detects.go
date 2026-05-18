package analyzer

import (
	"T1-DNSAnalysis/models"
)

func DetectBlocking(
	results []models.DNSResponse,
) bool {

	for _, r := range results {

		if r.RCode == -1 {
			continue
		}

		if r.RCode == 3 {
			return true
		}

		if r.RCode == 5 {
			return true
		}

		/*
			IPs clássicos
			de bloqueio
		*/

		for _, ip := range r.IPs {

			if ip == "127.0.0.1" ||
				ip == "0.0.0.0" {

				return true
			}
		}
	}

	return false
}

func DetectConsensus(
	results []models.DNSResponse,
) models.ConsensusResponse {

	ipCount := make(
		map[string]int,
	)

	response := models.ConsensusResponse{
		Outliers: []models.Outlier{},
	}

	for _, r := range results {

		if r.RCode == -1 {
			continue
		}

		for _, ip := range r.IPs {

			ipCount[ip]++
		}
	}

	// Identifica consenso
	var major string

	max := 0

	for ip, n := range ipCount {

		if n > max {

			max = n

			major = ip
		}
	}

	response.Consensus =
		major

	//	Identifica outliers
	for _, r := range results {

		if r.RCode == -1 {
			continue
		}

		for _, ip := range r.IPs {

			if ip != major {

				response.Outliers =
					append(
						response.Outliers,

						models.Outlier{

							Server: r.Server,

							IP: ip,
						},
					)
			}
		}
	}

	return response
}
