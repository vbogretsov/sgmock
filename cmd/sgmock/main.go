package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/labstack/echo/middleware"

	"github.com/vbogretsov/sgmock"
	"github.com/vbogretsov/sgmock/api"
)

const (
	name    = "sgmock"
	usage   = "SendGrid API v3 mock"
	version = "0.0.1"

	portHelp = "service listen port"
	keyHelp  = "test sendgrid API key, used for send request authorization"
)

type argT struct {
	port *int
	key  *string
}

var (
	args   = argT{}
	parser = argparse.NewParser(fmt.Sprintf("%s %s", name, version), usage)
)

func init() {
	args.port = parser.Int(
		"",
		"port",
		&argparse.Options{
			Required: false,
			Default:  9001,
			Help:     portHelp,
		})
	args.key = parser.String(
		"",
		"key",
		&argparse.Options{
			Required: true,
			Help:     portHelp,
		})
}

func run() error {
	if err := parser.Parse(os.Args); err != nil {
		return err
	}

	e := api.New(*args.key, sgmock.New())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if err := e.Start(fmt.Sprintf(":%d", *args.port)); err != nil {
		e.Logger.Fatal(err)
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
