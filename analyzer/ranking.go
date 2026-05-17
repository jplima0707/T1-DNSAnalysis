package analyzer

import (
	"T1-DNSAnalysis/models"
	"sort"
)

func SortRanking(

	data []models.
		BenchmarkResult,

) {

	sort.Slice(

		data,

		func(i, j int) bool {

			return data[i].
				Avg <
				data[j].
					Avg
		},
	)
}
