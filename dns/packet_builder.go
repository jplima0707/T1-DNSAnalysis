package dns

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"T1-DNSAnalysis/utils"
)

func BuildDNSQuery(domain string, id uint16) ([]byte, error) {
	normalized, err := utils.NormalizeDomain(domain)
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)

	fields := []uint16{
		id,
		0x0100,
		1,
		0,
		0,
		0,
	}

	for _, field := range fields {
		if err := binary.Write(buffer, binary.BigEndian, field); err != nil {
			return nil, fmt.Errorf("falha ao escrever cabecalho DNS: %w", err)
		}
	}

	if _, err := buffer.Write(utils.EncodeDomain(normalized)); err != nil {
		return nil, fmt.Errorf("falha ao escrever QNAME: %w", err)
	}

	if err := binary.Write(buffer, binary.BigEndian, uint16(1)); err != nil {
		return nil, fmt.Errorf("falha ao escrever QTYPE: %w", err)
	}

	if err := binary.Write(buffer, binary.BigEndian, uint16(1)); err != nil {
		return nil, fmt.Errorf("falha ao escrever QCLASS: %w", err)
	}

	return buffer.Bytes(), nil
}
