package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Retorna uma aplicação a ser rodada no terminal
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de comando"
	app.Usage = "Realiza pesquisas de Ip e Servidores na rede mundial de computadores"

	flags := []cli.Flag {
		cli.StringFlag{
			Name: "host",
			Value: "",
		},
	}

	app.Commands = [] cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPs de sites/servidores na Internet",
			Flags: flags,
			Action: buscarIps,
		},{
			Name: "servidores",
			Usage: "Buscar Servidores na Internet",
			Flags: flags,
			Action: buscarServidor,
		},
	}

	return app
}

func buscarIps(c *cli.Context)  {
	host := c.String("host")
	
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidor(c*cli.Context)  {
	host := c.String("host")

	servidores, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, servidor := range servidores{
		fmt.Println(servidor)
	}
}