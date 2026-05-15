package dns

import (
	"T1-DNSAnalysis/models"
	"sync"
)

func QueryServer(
	server string,
	domain string,
) models.DNSResponse {

	packet := BuildDNSQuery(
		domain,
	)

	response,
		elapsed,
		err := SendUDPQuery(
		server,
		packet,
	)

	if err != nil {

		return models.DNSResponse{
			Server: server,
			Error:  err,
		}
	}

	ips, rcode := ParseResponse(
		response,
	)

	return models.DNSResponse{
		Server:       server,
		IPs:          ips,
		RCode:        rcode,
		ResponseTime: elapsed,
	}
}

func SequentialScan(

	servers []string,
	domain string,

) []models.DNSResponse {

	var results []models.DNSResponse

	for _, server := range servers {

		result := QueryServer(
			server,
			domain,
		)

		results = append(
			results,
			result,
		)

	}

	return results

}

func ConcurrentScan(

	servers []string,
	domain string,

) []models.DNSResponse {

	var wg sync.WaitGroup

	results := make(
		chan models.DNSResponse,
		len(servers),
	)

	for _, server := range servers {

		wg.Add(1)

		go func(
			s string,
		) {

			defer wg.Done()

			result := QueryServer(
				s,
				domain,
			)

			results <- result

		}(server)

	}

	wg.Wait()

	close(results)

	var output []models.DNSResponse

	for r := range results {

		output = append(
			output,
			r,
		)

	}

	return output

}
