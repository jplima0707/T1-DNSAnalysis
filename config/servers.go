package config

type DNSServer struct {
	Name string
	IP   string
}

var Servers = []DNSServer{
	{"Google", "8.8.8.8"},
	{"Google Secondary", "8.8.4.4"},
	{"Cloudflare", "1.1.1.1"},
	{"Cloudflare Secondary", "1.0.0.1"},
	{"Quad9", "9.9.9.9"},
}

func GetIPs() []string {

	var ips []string

	for _, s := range Servers {

		ips = append(
			ips,
			s.IP,
		)
	}

	return ips

}
