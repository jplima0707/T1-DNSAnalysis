package dns

import (
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"

	"T1-DNSAnalysis/config"
	"T1-DNSAnalysis/models"
)

func SendDoTQuery(server config.DNSServer, domain string, timeout time.Duration) (*models.ParsedDNSResponse, time.Duration, error) {
	id := uint16(rand.Uint32())
	packet, err := BuildDNSQuery(domain, id)
	if err != nil {
		return nil, 0, err
	}

	dialer := &net.Dialer{Timeout: timeout}
	conn, err := tls.DialWithDialer(dialer, "tcp", net.JoinHostPort(server.IP, "853"), &tls.Config{
		ServerName: server.DoTHost,
		MinVersion: tls.VersionTLS12,
	})
	if err != nil {
		return nil, 0, err
	}
	defer conn.Close()

	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		return nil, 0, err
	}

	framed := make([]byte, 2+len(packet))
	binary.BigEndian.PutUint16(framed[:2], uint16(len(packet)))
	copy(framed[2:], packet)

	start := time.Now()

	if _, err := conn.Write(framed); err != nil {
		return nil, 0, err
	}

	lengthPrefix := make([]byte, 2)
	if _, err := io.ReadFull(conn, lengthPrefix); err != nil {
		return nil, time.Since(start), err
	}

	responseLength := int(binary.BigEndian.Uint16(lengthPrefix))
	if responseLength <= 0 {
		return nil, time.Since(start), fmt.Errorf("resposta DoT com tamanho invalido")
	}

	response := make([]byte, responseLength)
	if _, err := io.ReadFull(conn, response); err != nil {
		return nil, time.Since(start), err
	}

	elapsed := time.Since(start)
	parsed, err := ParseResponse(response)
	if err != nil {
		return nil, elapsed, err
	}

	if parsed.ID != id {
		return nil, elapsed, fmt.Errorf("ID da resposta nao corresponde ao da consulta")
	}

	return parsed, elapsed, nil
}
