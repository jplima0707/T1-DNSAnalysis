package dns

import (
	"T1-DNSAnalysis/models"
	"T1-DNSAnalysis/utils"
	"sync"
	"time"
)

func QueryServer(
	server models.DNSServer,
	domain string,
	protocol models.Protocol,

) models.DNSResponse {

	packet := BuildDNSQuery(
		domain,
	)

	var (
		response []byte
		elapsed  time.Duration
		err      error
	)

	switch protocol {

	case models.UDP:

		response,
			elapsed,
			err =

			SendUDPQuery(
				server.IP,
				packet,
			)

	case models.DOT:

		response,
			elapsed,
			err =

			SendDoTQuery(
				server.IP,
				packet,
			)

	}

	if err != nil {

		return models.DNSResponse{
			Server: server.IP,
			Error:  err,
			RCode:  utils.ExplainRcode(-1),
		}
	}

	ips,
		rcode := ParseResponse(
		response,
	)

	return models.DNSResponse{
		Server:       server.IP,
		IPs:          ips,
		RCode:        utils.ExplainRcode(rcode),
		ResponseTime: elapsed,
	}
}

func SequentialScan(
	servers []models.DNSServer,
	domain string,
	protocol models.Protocol,

) []models.DNSResponse {

	var results []models.DNSResponse

	for _, server := range servers {

		result := QueryServer(
			server,
			domain,
			protocol,
		)

		results = append(
			results,
			result,
		)

	}

	return results

}

func ConcurrentScan(
	servers []models.DNSServer,
	domain string,
	protocol models.Protocol,

) []models.DNSResponse {

	var wg sync.WaitGroup

	results := make(
		chan models.DNSResponse,
		len(servers),
	)

	for _, server := range servers {

		wg.Add(1)

		go func(
			s models.DNSServer,
		) {

			defer wg.Done()

			result := QueryServer(
				s,
				domain,
				protocol,
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
