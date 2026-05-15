package main

import (
	"T1-DNSAnalysis/config"
	"T1-DNSAnalysis/dns"
	"T1-DNSAnalysis/models"
	"T1-DNSAnalysis/utils"
	"fmt"
	"time"
)

func main() {

	domain := "reddit.com"

	servers := config.GetIPs()

	fmt.Println(
		"=========== SCAN SEQUENCIAL ==========",
	)

	start := time.Now()

	seq := dns.SequentialScan(

		servers,
		domain,
	)

	fmt.Println(
		"Tempo total:",
		time.Since(start),
	)

	printResults(seq)

	fmt.Println()

	fmt.Println(
		"=========== SCAN GOROUTINES ============",
	)

	start = time.Now()

	conc := dns.ConcurrentScan(

		servers,
		domain,
	)

	fmt.Println(
		"Tempo total:",
		time.Since(start),
	)

	printResults(conc)

}

func printResults(

	results []models.DNSResponse,

) {

	for _, r := range results {

		fmt.Println(
			"Servidor:",
			r.Server,
		)

		if r.Error != nil {

			fmt.Println(
				r.Error,
			)

			continue
		}

		fmt.Println(
			"Tempo:",
			r.ResponseTime,
		)

		fmt.Println(
			"IPs:",
			r.IPs,
		)

		fmt.Println(
			"RCode:",
			utils.ExplainRcode(
				r.RCode,
			),
		)

		fmt.Println()
	}

}
