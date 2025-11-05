package main

import (
	"base_app/config"
	"base_app/logger"
	"base_app/server"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	commandArgs := os.Args
	config.SetConfigFileFromArgs(commandArgs)
	config.Load(commandArgs)

	logger.SetupLogger()

	app := &cli.App{
		Name:  "Base_App",
		Usage: "This service contains the apis",
		Commands: []*cli.Command{
			{
				Name:   "api",
				Usage:  "run api server",
				Action: server.StartAPI,
			},
		},
	}

	if e := app.Run(os.Args); e != nil {
		fmt.Println(e)
	}
}
