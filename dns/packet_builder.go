package dns

import (
	"T1-DNSAnalysis/utils"
	"bytes"
	"encoding/binary"
	"math/rand"
)

func BuildDNSQuery(
	domain string,
) []byte {

	buffer := new(bytes.Buffer)

	id := uint16(rand.Intn(65535))

	binary.Write(
		buffer,
		binary.BigEndian,
		id,
	)

	binary.Write(
		buffer,
		binary.BigEndian,
		uint16(0x0100),
	)

	binary.Write(
		buffer,
		binary.BigEndian,
		uint16(1),
	)

	binary.Write(
		buffer,
		binary.BigEndian,
		uint16(0),
	)

	binary.Write(
		buffer,
		binary.BigEndian,
		uint16(0),
	)

	binary.Write(
		buffer,
		binary.BigEndian,
		uint16(0),
	)

	buffer.Write(
		utils.EncodeDomain(domain),
	)

	binary.Write(
		buffer,
		binary.BigEndian,
		uint16(1),
	)

	binary.Write(
		buffer,
		binary.BigEndian,
		uint16(1),
	)

	return buffer.Bytes()

}
