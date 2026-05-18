package main

import (
	"T1-DNSAnalysis/analyzer"
	"T1-DNSAnalysis/config"
	"T1-DNSAnalysis/dns"
	"T1-DNSAnalysis/models"
	"T1-DNSAnalysis/utils"
	"fmt"
)

func main() {

	domain := "reddit.com"

	// UDP
	runUDP(domain)

	// DoT
	runDoT(domain)

}

func runUDP(domain string) {

	fmt.Println(
		"\n===================================",
	)

	fmt.Println(
		"DNS UDP",
	)

	fmt.Println(
		"===================================",
	)

	results := dns.BenchmarkAllServers(

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

func runDoT(domain string) {

	fmt.Println(
		"\n===================================",
	)

	fmt.Println(
		"DNS over TLS",
	)

	fmt.Println(
		"===================================",
	)

	results := dns.BenchmarkAllServers(

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

		if len(
			r.Results,
		) > 0 {

			first := r.Results[0]

			fmt.Println(
				"RCODE:",
				first.RCode,
			)

			fmt.Println(
				"IPs:",
				first.IPs,
			)
		}

		fmt.Println()
	}
}
