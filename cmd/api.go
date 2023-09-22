package cmd

import (
	"fmt"
	"joubertredrat/bexs-dev-test-2k23/internal/infra"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

func getApiCommand() *cli.Command {
	return &cli.Command{
		Name:    "api",
		Aliases: []string{},
		Usage:   "Open HTTP api to listen",
		Action: func(c *cli.Context) error {
			config, err := infra.NewConfig()
			if err != nil {
				return err
			}

			r := gin.Default()
			if err := r.SetTrustedProxies(nil); err != nil {
				return err
			}

			apiBaseController := infra.NewApiBaseController()

			r.NoRoute(apiBaseController.HandleNotFound)

			return r.Run(fmt.Sprintf("%s:%s", config.ApiHost, config.ApiPort))
		},
	}
}
