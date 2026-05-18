package config

import "T1-DNSAnalysis/models"

var Servers = []models.DNSServer{

	//Sem filtro
	{Name: "Google", IP: "8.8.8.8"},
	{Name: "Google Secondary", IP: "8.8.4.4"},

	{Name: "Cloudflare", IP: "1.1.1.1"},
	{Name: "Cloudflare Secondary", IP: "1.0.0.1"},

	{Name: "Quad9 No Filter", IP: "9.9.9.10"},

	{Name: "Verisign", IP: "64.6.64.6"},

	//Segurança
	{Name: "Quad9 Secure", IP: "9.9.9.9"},

	{Name: "OpenDNS", IP: "208.67.222.222"},

	{Name: "CleanBrowsing Security",
		IP: "185.228.168.9"},

	{Name: "AdGuard DNS",
		IP: "94.140.14.14"},

	//Family
	{Name: "Cloudflare Family",
		IP: "1.1.1.3"},

	{Name: "OpenDNS Family",
		IP: "208.67.222.123"},

	{Name: "CleanBrowsing Family",
		IP: "185.228.168.168"},

	{Name: "AdGuard Family",
		IP: "94.140.14.15"},

	//Adicionados
	{Name: "ControlD", IP: "76.76.2.0"},

	{Name: "Comodo Secure", IP: "8.26.56.26"},

	{Name: "Comodo Secure Secondary",
		IP: "8.20.247.20"},

	{Name: "Alternate DNS",
		IP: "76.76.19.19"},

	{Name: "DNS Sistema Operacional",
		IP: "192.168.0.1"},
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
		{Name: "Google", IP: "dns.google"},
		{Name: "Cloudflare", IP: "one.one.one.one"},
		{Name: "Quad9", IP: "dns.quad9.net"},
	}
}
