package dns

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"T1-DNSAnalysis/models"
)

func SendUDPQuery(serverIP string, domain string, timeout time.Duration) (*models.ParsedDNSResponse, time.Duration, error) {
	id := uint16(rand.Uint32())
	packet, err := BuildDNSQuery(domain, id)
	if err != nil {
		return nil, 0, err
	}

	conn, err := net.Dial("udp", net.JoinHostPort(serverIP, "53"))
	if err != nil {
		return nil, 0, err
	}
	defer conn.Close()

	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		return nil, 0, err
	}

	start := time.Now()

	if _, err := conn.Write(packet); err != nil {
		return nil, 0, err
	}

	response := make([]byte, 512)
	n, err := conn.Read(response)
	elapsed := time.Since(start)
	if err != nil {
		return nil, elapsed, err
	}

	parsed, err := ParseResponse(response[:n])
	if err != nil {
		return nil, elapsed, err
	}

	if parsed.ID != id {
		return nil, elapsed, fmt.Errorf("ID da resposta nao corresponde ao da consulta")
	}

	return parsed, elapsed, nil
}
