package main

import (
	"fmt"
	"os"
	"log"

	"github.com/urfave/cli"
	"github.com/miekg/dns"

	"github.com/velp/go-rknresolver"
)


func main() {
	var dumpFile string
	var jsonFile string
	var countWorkers int
	var partSize int
	var dnsServer string

	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")

	app := cli.NewApp()
	app.Name = "rknresolver"
	app.Version = "0.0.1"
	app.Usage = "cli Roscomnadzor's dump.xml resolver and converter to JSON"

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "dump, d",
			Value: "",
			Usage: "path to XML dump file",
			Destination: &dumpFile,
		},
		cli.StringFlag{
			Name: "out, o",
			Value: "",
			Usage: "path to output JSON file",
			Destination: &jsonFile,
		},
		cli.IntFlag{
			Name: "workers, w",
			Value: 100,
			Usage: "count workers in resolver",
			Destination: &countWorkers,
		},
		cli.IntFlag{
			Name: "part, p",
			Value: 1000,
			Usage: "part size",
			Destination: &partSize,
		},
		cli.StringFlag{
			Name: "dns, n",
			Value: config.Servers[0]+":"+config.Port,
			Usage: "DNS server for resolving",
			Destination: &dnsServer,
		},
	}

	app.Action = func(c *cli.Context) error {
		if dumpFile == "" {
			return fmt.Errorf(`Bad argument "dump"`)
		}
		if jsonFile == "" {
			return fmt.Errorf(`Bad argument "out"`)
		}
		dump, err := rknresolver.Parse(dumpFile)
		if err != nil {
			return err
		}
		logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)
		resolver := rknresolver.NewResolver(countWorkers, partSize, dnsServer, logger)
		dump = resolver.ActualizeReg(dump)
		dump.WriteJSONFile(jsonFile)
		return nil
	}

	app.Run(os.Args)
}