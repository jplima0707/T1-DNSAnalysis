package utils

import "strings"

func EncodeDomain(domain string) []byte {

	parts := strings.Split(domain, ".")

	var result []byte

	for _, part := range parts {

		result = append(
			result,
			byte(len(part)),
		)

		result = append(
			result,
			[]byte(part)...,
		)
	}

	result = append(result, 0)

	return result
}
