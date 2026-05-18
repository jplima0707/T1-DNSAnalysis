package utils

func ExplainRcode(
	code int,
) string {

	switch code {

	case -1:
		return "NETWORK ERROR"
	case 0:
		return "OK"

	case 3:
		return "NXDOMAIN"

	case 5:
		return "REFUSED"

	default:
		return "UNKNOWN"

	}

}
