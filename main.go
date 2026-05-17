package main

import (
	"fmt"

	"T1-DNSAnalysis/analyzer"
	"T1-DNSAnalysis/config"
	"T1-DNSAnalysis/dns"
)

func main() {

	domain := "www.example.com"

	results := dns.
		BenchmarkAllServers(

			config.GetIPs(),

			domain,

			10,
		)

	analyzer.
		SortRanking(
			results,
		)

	fmt.Println(
		"========== RANKING ===========",
	)

	for i, r := range results {

		fmt.Printf(

			"%d - %s\n",

			i+1,

			r.Server,
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
			"%\n",
		)

	}

}
