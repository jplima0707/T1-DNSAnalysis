package dns

import (
	"net"
	"time"
)

func SendUDPQuery(
	server string,
	packet []byte,
) ([]byte, time.Duration, error) {

	conn, err := net.Dial(
		"udp",
		server+":53",
	)

	if err != nil {
		return nil, 0, err
	}

	defer conn.Close()

	conn.SetDeadline(
		time.Now().Add(
			5 * time.Second,
		),
	)

	start := time.Now()

	_, err = conn.Write(packet)

	if err != nil {
		return nil, 0, err
	}

	response := make(
		[]byte,
		512,
	)

	n, err := conn.Read(
		response,
	)

	elapsed := time.Since(start)

	if err != nil {
		return nil, 0, err
	}

	return response[:n], elapsed, nil

}
