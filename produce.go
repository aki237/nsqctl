package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bitly/go-nsq"
	cli "gopkg.in/urfave/cli.v2"
)

func produce(ctx *cli.Context) error {
	topic := ctx.String("topic")
	delay := ctx.String("delay")
	address := ctx.String("address")
	ddu := getDuration(delay)

	message := ctx.Args().First()
	if message == "" {
		return errors.New("Cannot publish an empty message")
	}

	config := nsq.NewConfig()
	w, err := nsq.NewProducer(address, config)
	if err != nil {
		return err
	}

	defer w.Stop()
	if ddu.Nanoseconds() != 0 {
		return w.DeferredPublish(topic, ddu, []byte(message))
	}
	return w.Publish(topic, []byte(message))
}

var suffixes = map[string]time.Duration{
	"ms": time.Millisecond,
	"s":  time.Second,
	"m":  time.Minute,
	"h":  time.Hour,
	"d":  24 * time.Hour,
	"w":  7 * 24 * time.Hour,
	"M":  30 * 7 * 24 * time.Hour,
	"y":  365 * 30 * 7 * 24 * time.Hour,
}

func getDuration(raw string) time.Duration {
	rxp := regexp.MustCompilePOSIX("^[0-9]+(ms|s|m|h|d|w|M|y)$")

	if !rxp.MatchString(raw) {
		fmt.Printf("[warn] Invalid delay: %s\n", raw)
		return 0 * time.Second
	}

	var num time.Duration

	for key, val := range suffixes {
		if strings.HasSuffix(raw, key) {
			n, _ := strconv.Atoi(strings.TrimSuffix(raw, key))
			num = time.Duration(n) * val
			break
		}
	}

	return num
}
