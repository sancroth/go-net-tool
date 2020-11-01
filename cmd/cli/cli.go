package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "DNS Records CLI"
	app.Usage = "Allows user to query IPs, CNAMEs, MX records"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "nps",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Look up the Name server for a given Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Look up the IP for a given Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Look up the MX records for a given Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Look up the CNAME records for a given Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cn, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					return err
				}
				fmt.Println(cn)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
