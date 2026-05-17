package utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ExplainRcode(code int) string {
	switch code {
	case 0:
		return "NOERROR"
	case 1:
		return "FORMERR"
	case 2:
		return "SERVFAIL"
	case 3:
		return "NXDOMAIN"
	case 4:
		return "NOTIMP"
	case 5:
		return "REFUSED"
	default:
		return "RCODE-" + strconv.Itoa(code)
	}
}

func FormatRCodes(values map[int]int) string {
	if len(values) == 0 {
		return "nenhum"
	}

	var keys []int
	for key := range values {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	parts := make([]string, 0, len(keys))
	for _, key := range keys {
		parts = append(parts, fmt.Sprintf("%s=%d", ExplainRcode(key), values[key]))
	}

	return strings.Join(parts, ", ")
}

func FormatIPs(ips []string) string {
	if len(ips) == 0 {
		return "nenhum"
	}
	return strings.Join(ips, ", ")
}

func JoinFindings(findings []string) string {
	if len(findings) == 0 {
		return "nenhum"
	}
	return strings.Join(findings, " | ")
}

func SortedKeys(values map[string]struct{}) []string {
	output := make([]string, 0, len(values))
	for value := range values {
		output = append(output, value)
	}
	sort.Strings(output)
	return output
}

func JoinIPs(ips []string) string {
	return strings.Join(ips, ",")
}

func Itoa(value int) string {
	return strconv.Itoa(value)
}
