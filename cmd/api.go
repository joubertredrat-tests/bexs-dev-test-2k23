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
			paymentRepository := infra.NewPaymentRepositoryMysql(db)

			exchangeStatic := infra.NewExchangeStatic(config.RateUsd, config.RateEur, config.RateGbp)

			usecaseCreatePartner := application.NewUsecaseCreatePartner(partnerRepository)
			usecaseCreatePayment := application.NewUsecaseCreatePayment(
				partnerRepository,
				paymentRepository,
				exchangeStatic,
				config.PaymentDuplicatedSeconds,
			)

			apiBaseController := infra.NewApiBaseController()
			partnerController := infra.NewPartnerController()
			paymentController := infra.NewPaymentController()

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
				rp := ra.Group("/payments")
				{
					rp.POST(
						"",
						infra.JSONBodyMiddleware(),
						paymentController.HandleCreate(usecaseCreatePayment),
					)
				}
			}

			return r.Run(fmt.Sprintf("%s:%s", config.ApiHost, config.ApiPort))
		},
	}
}
