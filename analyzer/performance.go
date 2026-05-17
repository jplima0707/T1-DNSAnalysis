package analyzer

import (
	"T1-DNSAnalysis/models"
	"time"
)

type Stats struct {
	Server string
	Avg    time.Duration
	Min    time.Duration
	Max    time.Duration
	Loss   float64
}

func ComputeStats(
	results []models.DNSResponse,
) Stats {

	var total time.Duration

	min := results[0].
		ResponseTime

	max := results[0].
		ResponseTime

	lost := 0

	for _, r := range results {

		if r.Error != nil {

			lost++

			continue
		}

		total += r.
			ResponseTime

		if r.ResponseTime < min {

			min = r.ResponseTime
		}

		if r.ResponseTime > max {

			max = r.ResponseTime
		}
	}

	return Stats{

		Server: results[0].Server,

		Avg: total /
			time.Duration(
				len(results)-lost,
			),

		Min: min,

		Max: max,

		Loss: float64(lost) /
			float64(len(results)) * 100,
	}
}
