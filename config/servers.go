package config

type DNSServer struct {
	Name      string
	Provider  string
	IP        string
	DoTHost   string
	IsPublic  bool
	IsExample bool
}

var Servers = []DNSServer{
	{Name: "Google Primary", Provider: "Google Public DNS", IP: "8.8.8.8", DoTHost: "dns.google", IsPublic: true},
	{Name: "Google Secondary", Provider: "Google Public DNS", IP: "8.8.4.4", DoTHost: "dns.google", IsPublic: true},
	{Name: "Cloudflare Primary", Provider: "Cloudflare", IP: "1.1.1.1", DoTHost: "cloudflare-dns.com", IsPublic: true},
	{Name: "Cloudflare Secondary", Provider: "Cloudflare", IP: "1.0.0.1", DoTHost: "cloudflare-dns.com", IsPublic: true},
	{Name: "Quad9 Primary", Provider: "Quad9", IP: "9.9.9.9", DoTHost: "dns.quad9.net", IsPublic: true},
	{Name: "Quad9 Secondary", Provider: "Quad9", IP: "149.112.112.112", DoTHost: "dns.quad9.net", IsPublic: true},
	{Name: "OpenDNS Primary", Provider: "OpenDNS", IP: "208.67.222.222", DoTHost: "dns.opendns.com", IsPublic: true},
	{Name: "OpenDNS Secondary", Provider: "OpenDNS", IP: "208.67.220.220", DoTHost: "dns.opendns.com", IsPublic: true},
	{Name: "Comcast Primary", Provider: "Comcast Xfinity", IP: "75.75.75.75", IsExample: true},
	{Name: "Comcast Secondary", Provider: "Comcast Xfinity", IP: "75.75.76.76", IsExample: true},
}

func GetServers() []DNSServer {
	servers := make([]DNSServer, len(Servers))
	copy(servers, Servers)
	return servers
}

func FilterDoTServers(servers []DNSServer) []DNSServer {
	var output []DNSServer
	for _, server := range servers {
		if server.DoTHost != "" {
			output = append(output, server)
		}
	}
	return output
}
