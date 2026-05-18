package utils

import (
	"fmt"
	"strings"
)

func NormalizeDomain(domain string) (string, error) {
	normalized := strings.TrimSpace(strings.TrimSuffix(domain, "."))
	if normalized == "" {
		return "", fmt.Errorf("dominio vazio")
	}

	parts := strings.Split(normalized, ".")
	for _, part := range parts {
		if part == "" {
			return "", fmt.Errorf("dominio invalido: label vazia")
		}
		if len(part) > 63 {
			return "", fmt.Errorf("dominio invalido: label maior que 63 bytes")
		}
	}

	return normalized, nil
}

func EncodeDomain(domain string) []byte {
	parts := strings.Split(domain, ".")
	var result []byte

	for _, part := range parts {
		result = append(result, byte(len(part)))
		result = append(result, []byte(part)...)
	}

	result = append(result, 0)
	return result
}
