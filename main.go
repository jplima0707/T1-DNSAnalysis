package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"T1-DNSAnalysis/config"
	"T1-DNSAnalysis/dns"
)

func main() {
	domainFlag := flag.String("domain", "", "Dominio a ser consultado")
	attemptsFlag := flag.Int("attempts", 10, "Numero de consultas por servidor")
	timeoutFlag := flag.Int("timeout-ms", 3000, "Timeout por consulta em milissegundos")
	flag.Parse()

	domain := strings.TrimSpace(*domainFlag)
	if domain == "" && flag.NArg() > 0 {
		domain = flag.Arg(0)
	}

	if domain == "" {
		fmt.Println("Uso: go run . -domain exemplo.com [-attempts 10] [-timeout-ms 3000]")
		os.Exit(1)
	}

	servers := config.GetServers()
	timeout := dns.TimeoutFromMillis(*timeoutFlag)

	fmt.Printf("Dominio analisado: %s\n", domain)
	fmt.Printf("Servidores configurados: %d\n", len(servers))
	fmt.Printf("Tentativas por servidor: %d\n\n", *attemptsFlag)

	udpReports := dns.BenchmarkAllServers(servers, domain, *attemptsFlag, dns.ProtocolUDP, timeout)
	dns.ApplyCrossServerChecks(udpReports)
	dns.SortReports(udpReports)

	fmt.Println("========== RANKING UDP/53 ==========")
	dns.PrintReports(udpReports)

	dotServers := config.FilterDoTServers(servers)
	if len(dotServers) == 0 {
		fmt.Println("\nNenhum servidor com DoT configurado.")
		return
	}

	dotReports := dns.BenchmarkAllServers(dotServers, domain, *attemptsFlag, dns.ProtocolDoT, timeout)
	dns.ApplyCrossServerChecks(dotReports)
	dns.SortReports(dotReports)

	fmt.Println("\n========== RANKING DoT/853 ==========")
	dns.PrintReports(dotReports)

	fmt.Println("\n========== COMPARACAO UDP x DoT ==========")
	dns.PrintProtocolComparison(udpReports, dotReports)
}
