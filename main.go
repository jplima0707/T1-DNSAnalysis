package main

import (
	"fmt"

	"T1-DNSAnalysis/analyzer"
	"T1-DNSAnalysis/config"
	"T1-DNSAnalysis/dns"
	"T1-DNSAnalysis/models"
	"T1-DNSAnalysis/utils"
)

func main() {

	domain := "reddit.com"

	runUDP(domain)

	runDoT(domain)
}

func runUDP(
	domain string,
) {

	fmt.Println(
		"\n===================================",
	)

	fmt.Println(
		"DNS UDP",
	)

	fmt.Println(
		"===================================",
	)

	/*
		SCANNER
	*/

	scan := dns.ScanServers(
		config.GetIPs(),
		domain,
		models.UDP,
	)

	printScan(scan)

	/*
		BENCHMARK
	*/

	results :=
		dns.BenchmarkAllServers(
			config.GetIPs(),
			domain,
			10,
			models.UDP,
		)

	analyzer.SortRanking(
		results,
	)

	printBenchmark(
		results,
	)

	err := utils.SaveCSV(
		"udp_results.csv",
		results,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"\nCSV UDP criado",
	)
}

func runDoT(
	domain string,
) {

	fmt.Println(
		"\n===================================",
	)
	fmt.Println(
		"DNS over TLS",
	)
	fmt.Println(
		"===================================",
	)

	scan :=
		dns.ScanServers(
			config.GetDoTHosts(),
			domain,
			models.DOT,
		)

	printScan(
		scan,
	)

	results :=
		dns.BenchmarkAllServers(
			config.GetDoTHosts(),
			domain,
			10,
			models.DOT,
		)

	analyzer.SortRanking(
		results,
	)

	printBenchmark(
		results,
	)

	err := utils.SaveCSV(
		"dot_results.csv",
		results,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"\nCSV DoT criado",
	)
}

func printScan(
	scan models.ScanResult,
) {

	fmt.Println(
		"\nANÁLISE\n",
	)

	fmt.Println(
		"Bloqueado:",
		scan.Blocked,
	)

	fmt.Println(
		"Consenso:",
		scan.Consensus.Consensus,
	)

	if len(
		scan.Consensus.Outliers,
	) > 0 {

		fmt.Println(
			"\nDivergentes:",
		)

		for _, o := range scan.Consensus.Outliers {

			fmt.Printf(
				"%s => %s\n",

				o.Server,

				o.IPs,
			)
		}
	}

	fmt.Println()
}

func printBenchmark(
	results []models.BenchmarkResult,
) {

	fmt.Println(
		"\nRANKING\n",
	)

	for i, r := range results {

		fmt.Printf(
			"%d - %s = %s\n",

			i+1,

			r.ServerName,

			r.ServerIP,
		)

		fmt.Println(
			"AVG:",
			r.Avg,
		)

		fmt.Println(
			"MIN:",
			r.Min,
		)

		fmt.Println(
			"MAX:",
			r.Max,
		)

		fmt.Println(
			"LOSS:",
			r.Loss,
			"%",
		)

		valid :=
			getFirstValid(
				r.Results,
			)

		if valid != nil {

			fmt.Println(
				"RCODE:",
				utils.ExplainRcode(
					valid.RCode,
				),
			)

			fmt.Println(
				"IPs:",
				valid.IPs,
			)

		} else {

			fmt.Println(
				"Todas consultas falharam",
			)
		}

		fmt.Println()
	}
}

func getFirstValid(
	results []models.DNSResponse,
) *models.DNSResponse {

	for i := range results {

		if results[i].Error == nil {

			return &results[i]
		}
	}

	return nil
}
