package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	/* fmt.Println("hello world")
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	} */
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPS, CNAME's"

	flagString := cli.StringFlag{
		Name:  "host",
		Value: "example.com",
	}
	myFlags := []cli.Flag{
		&flagString,
	}
	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Look up the Name servers for a particular host",
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
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
