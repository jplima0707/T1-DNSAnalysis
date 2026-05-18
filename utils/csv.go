package utils

import (
	"T1-DNSAnalysis/models"
	"encoding/csv"
	"os"
	"strconv"
)

func SaveCSV(

	filename string,

	results []models.BenchmarkResult,

) error {

	file, err := os.Create(
		filename,
	)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(
		file,
	)

	defer writer.Flush()

	writer.Write(
		[]string{

			"Servidor",

			"Media(ms)",

			"Min(ms)",

			"Max(ms)",

			"Perda(%)",

			"Consultas",
		},
	)

	for _, r := range results {

		writer.Write(
			[]string{

				r.ServerName,

				strconv.FormatInt(
					r.Avg.Milliseconds(),
					10,
				),

				strconv.FormatInt(
					r.Min.Milliseconds(),
					10,
				),

				strconv.FormatInt(
					r.Max.Milliseconds(),
					10,
				),

				strconv.FormatFloat(
					r.Loss,
					'f',
					2,
					64,
				),

				strconv.Itoa(
					len(
						r.Results,
					),
				),
			},
		)
	}

	return nil
}
