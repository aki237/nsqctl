package main

import (
	"log"
	"os"

	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{}

	app.Name = "nsqctl"
	app.Usage = "Control nsq from commandline"
	app.Commands = []*cli.Command{
		{
			Name:   "produce",
			Usage:  "Publish a message to a topic",
			Action: produce,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "topic",
					Aliases: []string{"t"},
					Usage:   "Set the topic to which the message is to be published",
					Value:   "test",
				},
				&cli.StringFlag{
					Name:    "delay",
					Aliases: []string{"d"},
					Usage:   "Set the delay for a message",
					Value:   "0s",
				},
				&cli.StringFlag{
					Name:    "address",
					Aliases: []string{"a"},
					Usage:   "Set the remote address of nsqd",
					Value:   "127.0.0.1:4150",
				},
			},
		},
		{
			Name:   "consume",
			Usage:  "Listen for messages from channels",
			Action: consume,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "topic",
					Aliases: []string{"t"},
					Usage:   "Set the topic to which the message is to be read from",
					Value:   "test",
				},
				&cli.StringFlag{
					Name:    "channel",
					Aliases: []string{"c"},
					Usage:   "Set the topic from which the message is to be read from",
					Value:   "test",
				},
				&cli.BoolFlag{
					Name:    "wait",
					Aliases: []string{"w"},
					Usage:   "Set whether to wait for messages indefinitely till a Ctrl-C signal",
					Value:   false,
				},
				&cli.StringFlag{
					Name:    "address",
					Aliases: []string{"a"},
					Usage:   "Set the remote address of nsqd",
					Value:   "127.0.0.1:4150",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
