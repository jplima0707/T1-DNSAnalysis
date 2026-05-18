package config

import "T1-DNSAnalysis/models"

var Servers = []models.DNSServer{

	//Sem filtro
	{"Google", "8.8.8.8"},
	{"Google Secondary", "8.8.4.4"},

	{"Cloudflare", "1.1.1.1"},
	{"Cloudflare Secondary", "1.0.0.1"},

	{"Quad9 No Filter", "9.9.9.10"},

	{"Verisign", "64.6.64.6"},

	//Segurança
	{"Quad9 Secure", "9.9.9.9"},

	{"OpenDNS", "208.67.222.222"},

	{"CleanBrowsing Security",
		"185.228.168.9"},

	{"AdGuard DNS",
		"94.140.14.14"},

	//Family
	{"Cloudflare Family",
		"1.1.1.3"},

	{"OpenDNS Family",
		"208.67.222.123"},

	{"CleanBrowsing Family",
		"185.228.168.168"},

	{"AdGuard Family",
		"94.140.14.15"},

	//Adicionados
	{"ControlD",
		"76.76.2.0"},

	{"Comodo Secure",
		"8.26.56.26"},

	{"Comodo Secure Secondary",
		"8.20.247.20"},

	{"Alternate DNS",
		"76.76.19.19"},

	{"DNS Sistema Operacional",
		"192.168.0.1"},
}

func GetIPs() []models.DNSServer {

	var ips []models.DNSServer

	for _, s := range Servers {

		ips = append(
			ips,
			s,
		)
	}

	return ips

}

func GetDoTHosts() []models.DNSServer {

	return []models.DNSServer{
		{"Google", "dns.google"},
		{"Cloudflare", "one.one.one.one"},
		{"Quad9", "dns.quad9.net"},
	}
}
