package dns

import (
	"encoding/binary"
	"net"
)

func skipDomainName(
	data []byte,
	offset int,
) int {

	for {

		length := int(data[offset])

		// fim do domínio
		if length == 0 {
			return offset + 1
		}

		// ponteiro DNS
		// ex: C0 0C
		if length&0xC0 == 0xC0 {
			return offset + 2
		}

		offset += length + 1
	}
}

func ParseResponse(
	data []byte,
) ([]string, int) {

	var ips []string

	// FLAGS bytes [2:4]
	flags := binary.BigEndian.Uint16(
		data[2:4],
	)

	// últimos 4 bits = RCODE
	rcode := int(
		flags & 0x000F,
	)

	// ANCOUNT bytes [6:8]
	answerCount := int(
		binary.BigEndian.Uint16(
			data[6:8],
		),
	)

	// HEADER DNS = 12 bytes
	offset := 12

	/*
	   QUESTION SECTION

	   nome domínio
	   TYPE
	   CLASS
	*/

	offset = skipDomainName(
		data,
		offset,
	)

	// TYPE + CLASS
	offset += 4

	/*
	   ANSWERS
	*/

	for i := 0; i < answerCount; i++ {

		// NAME
		offset = skipDomainName(
			data,
			offset,
		)

		// TYPE
		recordType := binary.BigEndian.Uint16(
			data[offset : offset+2],
		)

		offset += 2

		// CLASS
		offset += 2

		// TTL
		offset += 4

		// RDLENGTH
		rdLength := binary.BigEndian.Uint16(
			data[offset : offset+2],
		)

		offset += 2

		/*
		   TYPE 1 = A
		   IPv4 possui 4 bytes
		*/

		if recordType == 1 &&
			rdLength == 4 {

			ip := net.IP(
				data[offset : offset+4],
			)

			ips = append(
				ips,
				ip.String()+"\n",
			)
		}

		offset += int(
			rdLength,
		)
	}

	return ips, rcode
}
