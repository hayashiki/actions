package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	webhook string
	message string
)

func main() {
	log.Println("start")
	// TODO: read from
	fs := flag.NewFlagSet("pj bot", flag.ExitOnError)
	fs.StringVar(&webhook, "webhook", "", "webhook")
	fs.StringVar(&message, "message", "", "message")

	//path = "config.yml"
	if err := fs.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if webhook == "" || message == "" {
		fs.PrintDefaults()
		fmt.Fprintln(os.Stderr, errors.New("webhook or message is required"))
		os.Exit(1)
	}

	if err := Run(webhook, message); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Run(webhook, message string) error {
	return notify(webhook, message)
}
