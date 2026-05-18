package dns

import (
	"crypto/tls"
	"encoding/binary"
	"io"
	"time"
)

func SendDoTQuery(
	host string,
	packet []byte,
) ([]byte, time.Duration, error) {

	start := time.Now()

	conn, err := tls.Dial(

		"tcp",

		host+":853",

		&tls.Config{

			ServerName: host,
		},
	)

	if err != nil {

		return nil,
			0,
			err
	}

	defer conn.Close()

	conn.SetDeadline(

		time.Now().Add(
			5 * time.Second,
		),
	)

	/*
	   RFC7858:

	   2 bytes tamanho
	   + mensagem DNS
	*/

	message := make(

		[]byte,

		2+len(packet),
	)

	binary.BigEndian.
		PutUint16(

			message[0:2],

			uint16(
				len(packet),
			),
		)

	copy(

		message[2:],

		packet,
	)

	_, err =
		conn.Write(
			message,
		)

	if err != nil {

		return nil,
			0,
			err
	}

	/*
	  primeiros 2 bytes
	  dizem tamanho resposta
	*/

	sizeBuffer := make(
		[]byte,
		2,
	)

	_, err =
		io.ReadFull(

			conn,

			sizeBuffer,
		)

	if err != nil {

		return nil,
			0,
			err
	}

	responseSize :=

		binary.
			BigEndian.
			Uint16(
				sizeBuffer,
			)

	response := make(

		[]byte,

		responseSize,
	)

	_, err =
		io.ReadFull(

			conn,

			response,
		)

	elapsed := time.
		Since(start)

	if err != nil {

		return nil,
			0,
			err
	}

	return response,
		elapsed,
		nil

}
