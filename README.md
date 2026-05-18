# T1-DNSAnalysis

Ferramenta em Go para analisar respostas DNS tradicionais por UDP/53 e DNS over TLS por TCP/853, montando e interpretando os pacotes DNS manualmente, sem usar bibliotecas prontas de DNS.

## O que a aplicacao faz

- monta consultas DNS binarias manualmente conforme a RFC 1035
- envia consultas UDP para varios servidores DNS
- interpreta `RCODE` e extrai registros `A`
- executa multiplas consultas por servidor para calcular `avg`, `min`, `max` e `loss rate`
- gera ranking de desempenho
- detecta sinais de manipulacao, como `NXDOMAIN`, `REFUSED`, `0.0.0.0`, `127.0.0.1` e IPs divergentes
- executa consultas DoT usando `crypto/tls` com framing de 2 bytes conforme a RFC 7858
- imprime uma comparacao resumida entre UDP e DoT

## Estrutura

- [main.go](./main.go): CLI e execucao dos rankings
- [config/servers.go](./config/servers.go): servidores DNS publicos e exemplos de ISP
- [dns/packet_builder.go](./dns/packet_builder.go): montagem manual da query DNS
- [dns/packet_parser.go](./dns/packet_parser.go): parsing manual da resposta DNS
- [dns/udp_client.go](./dns/udp_client.go): cliente UDP/53
- [dns/tls_client.go](./dns/tls_client.go): cliente DoT/853
- [dns/scanner.go](./dns/scanner.go): benchmark, estatisticas e deteccao
- [dns/output.go](./dns/output.go): impressao do ranking e comparacao

## Como executar

O domínio usado pela aplicação é definido diretamente em `main.go` na variável `domain`. Basta alterar esse valor e executar:

```bash
go run .
```

Por exemplo, abra `main.go`, ajuste `domain := "internetbadguys.com"` para o domínio desejado e rode `go run .`.

## Como interpretar a saida

- `Sucessos/Perdas`: quantidade de respostas validas e consultas perdidas
- `Loss rate`: percentual de perdas no servidor
- `Avg/Min/Max`: tempos de resposta calculados apenas sobre consultas bem-sucedidas
- `RCODES`: distribuicao dos codigos de retorno DNS
- `IPs A`: conjunto de IPv4 observados para o dominio
- `Alertas`: sinais locais ou comparativos de bloqueio/manipulacao

## Captura no Wireshark ou tcpdump

### UDP tradicional

Filtro Wireshark:

```text
udp.port == 53
```

Filtro tcpdump:

```bash
tcpdump -i <interface> udp port 53
```

O que verificar:

- IP de origem: sua maquina
- IP de destino: resolvedor DNS
- porta de origem: efemera
- porta de destino: 53/UDP
- nome do dominio visivel em texto claro no payload DNS
- tamanho do pacote de consulta e resposta

### DNS over TLS

Filtro Wireshark:

```text
tcp.port == 853
```

Filtro tcpdump:

```bash
tcpdump -i <interface> tcp port 853
```

O que verificar:

- handshake TCP e TLS antes da consulta
- IP de destino na porta 853/TCP
- aumento na quantidade de pacotes em relacao ao UDP
- dominio nao visivel em texto claro dentro do payload por causa da criptografia TLS

## Observacoes importantes

- alguns servidores marcados como exemplo de ISP podem nao responder fora da rede deles
- servidores sem `DoTHost` configurado participam apenas do benchmark UDP
- o parser atual extrai registros `A`, que sao os exigidos no enunciado
