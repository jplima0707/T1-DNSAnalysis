package dns

import (
	"encoding/binary"
	"fmt"
	"net"

	"T1-DNSAnalysis/models"
)

func ParseResponse(data []byte) (*models.ParsedDNSResponse, error) {
	if len(data) < 12 {
		return nil, fmt.Errorf("resposta DNS menor que o cabecalho: %d bytes", len(data))
	}

	flags := binary.BigEndian.Uint16(data[2:4])
	questionCount := int(binary.BigEndian.Uint16(data[4:6]))
	answerCount := int(binary.BigEndian.Uint16(data[6:8]))

	parsed := &models.ParsedDNSResponse{
		ID:          binary.BigEndian.Uint16(data[0:2]),
		RCode:       int(flags & 0x000F),
		QuestionCnt: questionCount,
		AnswerCnt:   answerCount,
		Flags: models.DNSFlags{
			Response:           flags&0x8000 != 0,
			Authoritative:      flags&0x0400 != 0,
			Truncated:          flags&0x0200 != 0,
			RecursionDesired:   flags&0x0100 != 0,
			RecursionAvailable: flags&0x0080 != 0,
		},
	}

	offset := 12

	for i := 0; i < questionCount; i++ {
		var err error
		offset, err = skipDomainName(data, offset)
		if err != nil {
			return nil, err
		}

		if offset+4 > len(data) {
			return nil, fmt.Errorf("question section incompleta")
		}

		offset += 4
	}

	for i := 0; i < answerCount; i++ {
		var err error
		offset, err = skipDomainName(data, offset)
		if err != nil {
			return nil, err
		}

		if offset+10 > len(data) {
			return nil, fmt.Errorf("answer section incompleta")
		}

		recordType := binary.BigEndian.Uint16(data[offset : offset+2])
		offset += 2

		offset += 2

		ttl := binary.BigEndian.Uint32(data[offset : offset+4])
		offset += 4

		rdLength := int(binary.BigEndian.Uint16(data[offset : offset+2]))
		offset += 2

		if offset+rdLength > len(data) {
			return nil, fmt.Errorf("RDATA excede o tamanho do pacote")
		}

		if recordType == 1 && rdLength == 4 {
			ip := net.IP(data[offset : offset+rdLength]).String()
			parsed.Records = append(parsed.Records, models.DNSRecord{
				Type: recordType,
				TTL:  ttl,
				Data: ip,
			})
		}

		offset += rdLength
	}

	return parsed, nil
}

func skipDomainName(data []byte, offset int) (int, error) {
	for {
		if offset >= len(data) {
			return 0, fmt.Errorf("offset fora dos limites ao ler dominio")
		}

		length := int(data[offset])

		if length == 0 {
			return offset + 1, nil
		}

		if length&0xC0 == 0xC0 {
			if offset+1 >= len(data) {
				return 0, fmt.Errorf("ponteiro DNS incompleto")
			}
			return offset + 2, nil
		}

		offset++
		if offset+length > len(data) {
			return 0, fmt.Errorf("label DNS excede o tamanho do pacote")
		}
		offset += length
	}
}
