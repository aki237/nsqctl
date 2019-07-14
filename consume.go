package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/bitly/go-nsq"
	cli "gopkg.in/urfave/cli.v2"
)

func consume(ctx *cli.Context) error {
	topic := ctx.String("topic")
	channel := ctx.String("channel")
	address := ctx.String("address")
	wait := ctx.Bool("wait")

	if channel == "" {
		return errors.New("Cannot publish an empty message")
	}

	config := nsq.NewConfig()
	q, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return err
	}
	defer q.Stop()

	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Printf("Got a message: %s\n", message.Body)
		if !wait {
			q.StopChan <- 1
		}
		return nil
	}))

	err = q.ConnectToNSQD(address)
	if err != nil {
		return err
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	select {
	case <-q.StopChan:
	case <-sig:
	}

	fmt.Println("Done watching")

	return nil
}
