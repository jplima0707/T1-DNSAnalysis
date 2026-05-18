package dns

import (
	"T1-DNSAnalysis/models"
	"sync"
	"time"
)

func BenchmarkServer(
	server models.DNSServer,
	domain string,
	times int,
	protocol models.Protocol,

) models.BenchmarkResult {

	var results []models.DNSResponse

	for i := 0; i < times; i++ {

		r := QueryServer(
			server,
			domain,
			protocol,
		)

		results = append(
			results,
			r,
		)

		time.Sleep(
			100 * time.Millisecond,
		)
	}

	return ComputeBenchmark(
		server,
		results,
	)

}

func ComputeBenchmark(
	server models.DNSServer,
	results []models.DNSResponse,

) models.BenchmarkResult {
	total := time.Duration(0)
	loss := 0
	min := time.Hour
	max := time.Duration(0)
	valid := 0

	for _, r := range results {

		if r.Error != nil {

			loss++

			continue
		}

		valid++

		total += r.ResponseTime

		if r.ResponseTime < min {

			min = r.ResponseTime
		}

		if r.ResponseTime > max {

			max = r.ResponseTime
		}

	}

	lossPercent := (float64(loss) /
		float64(len(results)) * 100)

	avg := time.Duration(0)

	if valid > 0 {
		avg =
			total /
				time.Duration(valid)
	}

	if valid == 0 {
		min = 0
		max = 0
	}

	return models.BenchmarkResult{

		ServerName: server.Name,
		ServerIP:   server.IP,
		Results:    results,
		Avg:        avg,
		Min:        min,
		Max:        max,
		Loss:       lossPercent,
	}

}

func BenchmarkAllServers(
	servers []models.DNSServer,
	domain string,
	times int,
	protocol models.Protocol,

) []models.BenchmarkResult {

	var wg sync.WaitGroup

	channel := make(
		chan models.BenchmarkResult,
		len(servers),
	)

	for _, server := range servers {

		wg.Add(1)

		go func(
			s models.DNSServer,
		) {

			defer wg.Done()

			result :=
				BenchmarkServer(
					s,
					domain,
					times,
					protocol,
				)

			channel <- result

		}(server)

	}

	wg.Wait()

	close(channel)

	var output []models.
		BenchmarkResult

	for r := range channel {

		output = append(
			output,
			r,
		)
	}

	return output
}
