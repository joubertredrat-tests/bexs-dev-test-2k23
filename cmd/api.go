package cmd

import (
	"fmt"
	"joubertredrat/bexs-dev-test-2k23/internal/application"
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

			db, err := infra.GetDatabaseConnection(infra.GetMysqlDSN(
				config.DatabaseHost,
				config.DatabasePort,
				config.DatabaseName,
				config.DatabaseUser,
				config.DatabasePassword,
			))
			if err != nil {
				return err
			}

			partnerRepository := infra.NewPartnerRepositoryMysql(db)

			usecaseCreatePartner := application.NewUsecaseCreatePartner(partnerRepository)

			apiBaseController := infra.NewApiBaseController()
			partnerController := infra.NewPartnerController()

			r.NoRoute(apiBaseController.HandleNotFound)

			ra := r.Group("/api")
			infra.RegisterCustomValidator()
			{
				ra.GET("/status", apiBaseController.HandleStatus)
				ra.POST(
					"/partners",
					infra.JSONBodyMiddleware(),
					partnerController.HandleCreate(usecaseCreatePartner),
				)
			}

			return r.Run(fmt.Sprintf("%s:%s", config.ApiHost, config.ApiPort))
		},
	}
}
