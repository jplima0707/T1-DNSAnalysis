package dns

import (
	"T1-DNSAnalysis/analyzer"
	"T1-DNSAnalysis/models"
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
			Server: server.Name,
			RCode:  -1,
			Error:  err,
		}
	}

	ips,
		rcode :=
		ParseResponse(
			response,
		)

	return models.DNSResponse{
		Server:       server.Name,
		IPs:          ips,
		RCode:        rcode,
		ResponseTime: elapsed,
	}
}

func ScanServers(
	servers []models.DNSServer,
	domain string,
	protocol models.Protocol,

) models.ScanResult {

	var wg sync.WaitGroup

	channel := make(
		chan models.DNSResponse,
		len(servers),
	)

	for _, server := range servers {

		wg.Add(1)

		go func(
			s models.DNSServer,
		) {

			defer wg.Done()

			result :=
				QueryServer(
					s,
					domain,
					protocol,
				)

			channel <- result

		}(server)
	}

	wg.Wait()

	close(channel)

	var responses []models.DNSResponse

	for r := range channel {

		responses =
			append(
				responses,
				r,
			)
	}

	blocked :=

		analyzer.
			DetectBlocking(
				responses,
			)

	consensus :=

		analyzer.
			DetectConsensus(
				responses,
			)

	return models.ScanResult{
		Responses: responses,
		Blocked:   blocked,
		Consensus: consensus,
	}
}
