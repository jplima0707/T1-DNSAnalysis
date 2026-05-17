package dns

import (
	"fmt"
	"time"

	"T1-DNSAnalysis/models"
	"T1-DNSAnalysis/utils"
)

func PrintReports(reports []models.ServerReport) {
	for index, report := range reports {
		fmt.Printf("%d. %s (%s - %s)\n", index+1, report.ServerName, report.Provider, report.ServerIP)
		fmt.Printf("   Protocolo: %s\n", report.Protocol)
		fmt.Printf("   Sucessos/Perdas: %d/%d\n", report.Successes, report.Losses)
		fmt.Printf("   Loss rate: %.2f%%\n", report.LossRate)
		fmt.Printf("   Avg/Min/Max: %s / %s / %s\n", formatDuration(report.Avg), formatDuration(report.Min), formatDuration(report.Max))
		fmt.Printf("   RCODES: %s\n", utils.FormatRCodes(report.RCodes))
		fmt.Printf("   IPs A: %s\n", utils.FormatIPs(report.IPs))

		if len(report.Findings) == 0 {
			fmt.Println("   Alertas: nenhum sinal claro de bloqueio/manipulacao")
		} else {
			fmt.Printf("   Alertas: %s\n", utils.JoinFindings(report.Findings))
		}

		fmt.Println()
	}
}

func PrintProtocolComparison(udpReports []models.ServerReport, dotReports []models.ServerReport) {
	udpAvg, udpLoss := protocolTotals(udpReports)
	dotAvg, dotLoss := protocolTotals(dotReports)

	fmt.Printf("UDP  -> media agregada: %s | perda media: %.2f%% | visibilidade: dominio aparece em texto claro no pacote DNS.\n", formatDuration(udpAvg), udpLoss)
	fmt.Printf("DoT  -> media agregada: %s | perda media: %.2f%% | visibilidade: payload DNS fica protegido dentro do TLS.\n", formatDuration(dotAvg), dotLoss)
	fmt.Println("Pacotes esperados no Wireshark:")
	fmt.Println("UDP  -> 1 consulta + 1 resposta por tentativa bem-sucedida")
	fmt.Println("DoT  -> handshake TCP/TLS + consulta/resposta encapsuladas em TLS")
}

func protocolTotals(reports []models.ServerReport) (time.Duration, float64) {
	var (
		totalAvg  time.Duration
		totalLoss float64
		count     int
	)

	for _, report := range reports {
		if report.Avg > 0 {
			totalAvg += report.Avg
			count++
		}
		totalLoss += report.LossRate
	}

	if count == 0 || len(reports) == 0 {
		return 0, 100
	}

	return totalAvg / time.Duration(count), totalLoss / float64(len(reports))
}

func formatDuration(duration time.Duration) string {
	if duration == 0 {
		return "n/a"
	}
	return duration.String()
}
