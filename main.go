package main

import (
	"fmt"
	"os"

	"github.com/etsxxx/nightmare-monkey/nightmare"

	_ "net/http/pprof"

	"github.com/urfave/cli"
)

var Flags = []cli.Flag{
	cli.BoolFlag{
		Name:  "execute",
		Usage: "You must specify this option to see nightmare. If you do not, then the monkey will be gentle.",
	},
	cli.Uint64Flag{
		Name:  "min-interval",
		Usage: "minimum interval between nightmare (sec)",
		Value: 3600,
	},
	cli.Uint64Flag{
		Name:  "max-interval",
		Usage: "max interval between nightmare (sec)",
		Value: 21600,
	},
	cli.StringFlag{
		Name:  "day, d",
		Usage: "specify the day of week to allow nightmare, like crontab. 0-7 (0 or 7 is Sun)",
		Value: "1-5",
	},
	cli.StringFlag{
		Name:  "time, t",
		Usage: "specify the time to allow nightmare. 'HH:MM-HH:MM' format.",
		Value: "11:00-18:00",
	},
	cli.UintFlag{
		Name:  "port, p",
		Usage: "specify the port on which the API listens for connections",
		Value: 8080,
	},
}

func Action(c *cli.Context) {
	interval := c.Uint64("max-interval") - c.Uint64("min-interval")
	if interval <= 0 {
		fmt.Printf("invalid interval.")
		return
	}

	app, err := nightmare.New(
		int(c.Uint("port")),
		interval,
		c.String("day"),
		c.String("time"),
	)
	if err != nil {
		os.Exit(1)
	}

	app.Dryrun = !c.Bool("execute")
	app.Play()
}

func main() {
	app := cli.NewApp()
	app.Name = "Nightmare Monkey"
	app.Version = fmt.Sprintf("%s (rev:%s)", version, gitcommit)
	app.Usage = Usage
	app.Flags = Flags
	app.Action = Action

	app.Run(os.Args)
}
