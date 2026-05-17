package dns

import (
	"sort"
	"sync"
	"time"

	"T1-DNSAnalysis/config"
	"T1-DNSAnalysis/models"
	"T1-DNSAnalysis/utils"
)

func TimeoutFromMillis(timeoutMs int) time.Duration {
	if timeoutMs <= 0 {
		timeoutMs = 3000
	}
	return time.Duration(timeoutMs) * time.Millisecond
}

func BenchmarkAllServers(servers []config.DNSServer, domain string, attempts int, protocol string, timeout time.Duration) []models.ServerReport {
	if attempts <= 0 {
		attempts = 10
	}

	results := make(chan models.ServerReport, len(servers))
	var wg sync.WaitGroup

	for _, server := range servers {
		wg.Add(1)

		go func(s config.DNSServer) {
			defer wg.Done()
			results <- BenchmarkServer(s, domain, attempts, protocol, timeout)
		}(server)
	}

	wg.Wait()
	close(results)

	var reports []models.ServerReport
	for report := range results {
		reports = append(reports, report)
	}

	return reports
}

func BenchmarkServer(server config.DNSServer, domain string, attempts int, protocol string, timeout time.Duration) models.ServerReport {
	report := models.ServerReport{
		ServerName: server.Name,
		Provider:   server.Provider,
		ServerIP:   server.IP,
		Protocol:   protocol,
		Attempts:   attempts,
		RCodes:     make(map[int]int),
	}

	var durations []time.Duration
	uniqueIPs := make(map[string]struct{})

	for attempt := 1; attempt <= attempts; attempt++ {
		query := models.QueryResult{
			Attempt:  attempt,
			Protocol: protocol,
		}

		var (
			response *models.ParsedDNSResponse
			elapsed  time.Duration
			err      error
		)

		switch protocol {
		case ProtocolDoT:
			response, elapsed, err = SendDoTQuery(server, domain, timeout)
		default:
			response, elapsed, err = SendUDPQuery(server.IP, domain, timeout)
		}

		query.ResponseTime = elapsed

		if err != nil {
			query.Error = err.Error()
			report.Losses++
			report.Results = append(report.Results, query)
			continue
		}

		report.Successes++
		report.RCodes[response.RCode]++
		query.Response = response
		report.Results = append(report.Results, query)
		durations = append(durations, elapsed)

		for _, record := range response.Records {
			if record.Type == 1 {
				uniqueIPs[record.Data] = struct{}{}
			}
		}
	}

	report.LossRate = float64(report.Losses) * 100 / float64(report.Attempts)
	report.IPs = utils.SortedKeys(uniqueIPs)
	report.Findings = append(report.Findings, evaluateStandaloneFindings(report)...)

	if len(durations) > 0 {
		report.Min, report.Max, report.Avg = durationStats(durations)
	}

	return report
}

func ApplyCrossServerChecks(reports []models.ServerReport) {
	ipOwners := make(map[string][]string)

	for _, report := range reports {
		for _, ip := range report.IPs {
			ipOwners[ip] = append(ipOwners[ip], report.ServerName)
		}
	}

	allIPSets := make(map[string]struct{})
	for _, report := range reports {
		if len(report.IPs) == 0 {
			continue
		}
		allIPSets[utils.JoinIPs(report.IPs)] = struct{}{}
	}

	divergent := len(allIPSets) > 1

	for i := range reports {
		if divergent {
			reports[i].Findings = appendUnique(reports[i].Findings, "IPs divergentes entre servidores")
		}

		for _, ip := range reports[i].IPs {
			if len(ipOwners[ip]) == 1 {
				reports[i].Findings = appendUnique(reports[i].Findings, "IP exclusivo encontrado apenas neste servidor: "+ip)
			}
		}
	}
}

func SortReports(reports []models.ServerReport) {
	sort.Slice(reports, func(i, j int) bool {
		if reports[i].LossRate != reports[j].LossRate {
			return reports[i].LossRate < reports[j].LossRate
		}

		if reports[i].Avg != reports[j].Avg {
			if reports[i].Avg == 0 {
				return false
			}
			if reports[j].Avg == 0 {
				return true
			}
			return reports[i].Avg < reports[j].Avg
		}

		return reports[i].ServerName < reports[j].ServerName
	})
}

func durationStats(durations []time.Duration) (time.Duration, time.Duration, time.Duration) {
	min := durations[0]
	max := durations[0]
	var total time.Duration

	for _, duration := range durations {
		if duration < min {
			min = duration
		}
		if duration > max {
			max = duration
		}
		total += duration
	}

	return min, max, total / time.Duration(len(durations))
}

func evaluateStandaloneFindings(report models.ServerReport) []string {
	var findings []string

	for rcode, count := range report.RCodes {
		switch rcode {
		case 3:
			findings = append(findings, "Respostas NXDOMAIN observadas: "+utils.Itoa(count))
		case 5:
			findings = append(findings, "Respostas REFUSED observadas: "+utils.Itoa(count))
		}
	}

	for _, ip := range report.IPs {
		if ip == "0.0.0.0" {
			findings = append(findings, "Possivel bloqueio por sinkhole: 0.0.0.0")
		}
		if ip == "127.0.0.1" {
			findings = append(findings, "Possivel redirecionamento local: 127.0.0.1")
		}
	}

	if len(report.IPs) == 0 && report.Successes > 0 {
		findings = append(findings, "Sem registros A nas respostas bem-sucedidas")
	}

	if report.LossRate > 0 {
		findings = append(findings, "Houve perda de consultas")
	}

	return findings
}

func appendUnique(values []string, value string) []string {
	for _, current := range values {
		if current == value {
			return values
		}
	}
	return append(values, value)
}
