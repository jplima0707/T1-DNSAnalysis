package analyzer

import (
	"T1-DNSAnalysis/models"
	"fmt"
)

func DetectBlocking(
	results []models.DNSResponse,
) {

	for _, r := range results {

		fmt.Println(
			"\nServidor:",
			r.Server,
		)

		switch r.RCode {

		case 3:
			fmt.Println(
				"Possível NXDOMAIN",
			)

		case 5:
			fmt.Println(
				"Consulta recusada",
			)
		}

		for _, ip := range r.IPs {

			if ip == "127.0.0.1" ||
				ip == "0.0.0.0" {

				fmt.Println(
					"Possível bloqueio:",
					ip,
				)
			}
		}
	}
}

func DetectConsensus(
	results []models.DNSResponse,
) {
	ipCount := make(
		map[string]int,
	)

	for _, r := range results {

		for _, ip := range r.IPs {

			ipCount[ip]++
		}
	}

	var major string
	max := 0

	for ip, n := range ipCount {

		if n > max {

			max = n

			major = ip
		}
	}

	fmt.Println(
		"\nIP consenso:",
		major,
	)

	for _, r := range results {

		for _, ip := range r.IPs {

			if ip != major {

				fmt.Println(
					r.Server,
					"IP divergente:",
					ip,
				)
			}
		}
	}
}
