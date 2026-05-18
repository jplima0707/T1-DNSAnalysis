package analyzer

import (
	"T1-DNSAnalysis/models"
	"sort"
	"strings"
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

	response := models.ConsensusResponse{}

	groupCount := make(
		map[string]int,
	)

	serverGroups := make(
		map[string][]string,
	)

	for _, r := range results {

		if r.RCode == -1 {

			continue
		}

		ips := append(
			[]string{},
			r.IPs...,
		)

		sort.Strings(
			ips,
		)

		key := strings.Join(
			ips,
			";",
		)

		groupCount[key]++

		serverGroups[r.Server] = ips
	}

	max := 0
	var major string

	for k, n := range groupCount {

		if n > max {

			max = n
			major = k
		}
	}

	response.Consensus = major

	for server, ips := range serverGroups {

		key := strings.Join(
			ips,
			";",
		)

		if key != major {
			response.Outliers =
				append(
					response.Outliers,
					models.Outlier{
						Server: server,
						IPs:    ips,
					})
		}
	}

	return response
}
