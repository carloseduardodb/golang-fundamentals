package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Vai retornar a aplicação de linha de comando pronto para ser executada
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host, H",
			Value: "www.google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca o IP de endereços na internet",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "servers",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: findServers,
		},
	}
	return app
}

func findIps(c *cli.Context) {
	host := c.String("host")
	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServers(c *cli.Context) {
	host := c.String("host")
	servers, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
