package application_test

import (
	"context"
	"errors"
	"joubertredrat/bexs-dev-test-2k23/internal/application"
	"joubertredrat/bexs-dev-test-2k23/internal/domain"
	"joubertredrat/bexs-dev-test-2k23/pkg/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUsecaseCreatePartner(t *testing.T) {
	tests := []struct {
		name              string
		partnerRepository func(ctrl *gomock.Controller) domain.PartnerRepository
		input             application.UsecaseCreatePartnerInput
		partnerExpected   domain.Partner
		errExpected       error
	}{
		{
			name: "Test create partner with success",
			partnerRepository: func(ctrl *gomock.Controller) domain.PartnerRepository {
				repository := mock.NewMockPartnerRepository(ctrl)

				repository.
					EXPECT().
					GetByID(gomock.Any(), gomock.Eq("10")).
					Return(domain.Partner{}, nil).
					Times(1)
				repository.
					EXPECT().
					GetByDocument(gomock.Any(), gomock.Eq("1284498339812/0001")).
					Return(domain.Partner{}, nil).
					Times(1)
				repository.
					EXPECT().
					Create(gomock.Any(), gomock.AssignableToTypeOf(domain.Partner{})).
					Return(domain.Partner{
						ID:          "10",
						TradingName: "International Ecommerce",
						Document:    "1284498339812/0001",
						Currency: domain.Currency{
							Value: "USD",
						},
					}, nil).
					Times(1)

				return repository
			},
			input: application.UsecaseCreatePartnerInput{
				ID:          "10",
				TradingName: "International Ecommerce",
				Document:    "1284498339812/0001",
				Currency:    "USD",
			},
			partnerExpected: domain.Partner{
				ID:          "10",
				TradingName: "International Ecommerce",
				Document:    "1284498339812/0001",
				Currency: domain.Currency{
					Value: "USD",
				},
			},
			errExpected: nil,
		},
		{
			name: "Test create partner with unknown error on get by id from repository",
			partnerRepository: func(ctrl *gomock.Controller) domain.PartnerRepository {
				repository := mock.NewMockPartnerRepository(ctrl)

				repository.
					EXPECT().
					GetByID(gomock.Any(), gomock.Eq("10")).
					Return(domain.Partner{}, errors.New("database gone")).
					Times(1)

				return repository
			},
			input: application.UsecaseCreatePartnerInput{
				ID:          "10",
				TradingName: "International Ecommerce",
				Document:    "1284498339812/0001",
				Currency:    "USD",
			},
			partnerExpected: domain.Partner{},
			errExpected:     errors.New("database gone"),
		},
		{
			name: "Test create partner with ID already exists in repository",
			partnerRepository: func(ctrl *gomock.Controller) domain.PartnerRepository {
				repository := mock.NewMockPartnerRepository(ctrl)

				repository.
					EXPECT().
					GetByID(gomock.Any(), gomock.Eq("10")).
					Return(domain.Partner{
						ID:          "10",
						TradingName: "International Ecommerce",
						Document:    "1284498339812/0001",
						Currency: domain.Currency{
							Value: "USD",
						},
					}, nil).
					Times(1)

				return repository
			},
			input: application.UsecaseCreatePartnerInput{
				ID:          "10",
				TradingName: "International Ecommerce",
				Document:    "1284498339812/0001",
				Currency:    "USD",
			},
			partnerExpected: domain.Partner{},
			errExpected:     application.NewErrPartnerAlreadyExists("ID", "10"),
		},
		{
			name: "Test create partner with unknown error on get by document from repository",
			partnerRepository: func(ctrl *gomock.Controller) domain.PartnerRepository {
				repository := mock.NewMockPartnerRepository(ctrl)

				repository.
					EXPECT().
					GetByID(gomock.Any(), gomock.Eq("10")).
					Return(domain.Partner{}, nil).
					Times(1)
				repository.
					EXPECT().
					GetByDocument(gomock.Any(), gomock.Eq("1284498339812/0001")).
					Return(domain.Partner{}, errors.New("database gone")).
					Times(1)

				return repository
			},
			input: application.UsecaseCreatePartnerInput{
				ID:          "10",
				TradingName: "International Ecommerce",
				Document:    "1284498339812/0001",
				Currency:    "USD",
			},
			partnerExpected: domain.Partner{},
			errExpected:     errors.New("database gone"),
		},
		{
			name: "Test create partner with document already exists in repository",
			partnerRepository: func(ctrl *gomock.Controller) domain.PartnerRepository {
				repository := mock.NewMockPartnerRepository(ctrl)

				repository.
					EXPECT().
					GetByID(gomock.Any(), gomock.Eq("10")).
					Return(domain.Partner{}, nil).
					Times(1)
				repository.
					EXPECT().
					GetByDocument(gomock.Any(), gomock.Eq("1284498339812/0001")).
					Return(domain.Partner{
						ID:          "10",
						TradingName: "International Ecommerce",
						Document:    "1284498339812/0001",
						Currency: domain.Currency{
							Value: "USD",
						},
					}, nil).
					Times(1)

				return repository
			},
			input: application.UsecaseCreatePartnerInput{
				ID:          "10",
				TradingName: "International Ecommerce",
				Document:    "1284498339812/0001",
				Currency:    "USD",
			},
			partnerExpected: domain.Partner{},
			errExpected:     application.NewErrPartnerAlreadyExists("document", "1284498339812/0001"),
		},
		{
			name: "Test create partner with unknown error on create from repository",
			partnerRepository: func(ctrl *gomock.Controller) domain.PartnerRepository {
				repository := mock.NewMockPartnerRepository(ctrl)

				repository.
					EXPECT().
					GetByID(gomock.Any(), gomock.Eq("10")).
					Return(domain.Partner{}, nil).
					Times(1)
				repository.
					EXPECT().
					GetByDocument(gomock.Any(), gomock.Eq("1284498339812/0001")).
					Return(domain.Partner{}, nil).
					Times(1)
				repository.
					EXPECT().
					Create(gomock.Any(), gomock.AssignableToTypeOf(domain.Partner{})).
					Return(domain.Partner{}, errors.New("database gone")).
					Times(1)

				return repository
			},
			input: application.UsecaseCreatePartnerInput{
				ID:          "10",
				TradingName: "International Ecommerce",
				Document:    "1284498339812/0001",
				Currency:    "USD",
			},
			partnerExpected: domain.Partner{},
			errExpected:     errors.New("database gone"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()
			ctrl := gomock.NewController(t)

			usecase := application.NewUsecaseCreatePartner(test.partnerRepository(ctrl))
			partnerGot, errGot := usecase.Execute(ctx, test.input)

			assert.Equal(t, test.partnerExpected, partnerGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}

func TestUsecaseCreatePayment(t *testing.T) {
	tests := []struct {
		name              string
		partnerRepository func(ctrl *gomock.Controller) domain.PartnerRepository
		paymentRepository func(ctrl *gomock.Controller) domain.PaymentRepository
		exchange          func(ctrl *gomock.Controller) domain.Exchange
		input             application.UsecaseCreatePaymentInput
		paymentExpected   domain.Payment
		errExpected       error
	}{
		{
			name: "Test create payment with success",
			partnerRepository: func(ctrl *gomock.Controller) domain.PartnerRepository {
				repository := mock.NewMockPartnerRepository(ctrl)

				repository.
					EXPECT().
					GetByID(gomock.Any(), gomock.Eq("10")).
					Return(domain.Partner{
						ID:          "10",
						TradingName: "International Ecommerce",
						Document:    "1284498339812/0001",
						Currency: domain.Currency{
							Value: "USD",
						},
					}, nil).
					Times(1)

				return repository
			},
			paymentRepository: func(ctrl *gomock.Controller) domain.PaymentRepository {
				repository := mock.NewMockPaymentRepository(ctrl)

				consumer, _ := domain.NewConsumer("Oliver Tsubasa", "30243434597")
				amount, _ := domain.NewAmount("99.05")
				foreignAmount, _ := domain.NewAmount("470.49")

				repository.
					EXPECT().
					Create(gomock.Any(), gomock.AssignableToTypeOf(domain.Payment{})).
					Return(domain.Payment{
						ID:            "01HAW44PR1XK7B027RSFE8SAAY",
						PartnerID:     "10",
						Consumer:      consumer,
						Amount:        amount,
						ForeignAmount: foreignAmount,
					}, nil).
					Times(1)

				return repository
			},
			exchange: func(ctrl *gomock.Controller) domain.Exchange {
				exchange := mock.NewMockExchange(ctrl)

				amount, _ := domain.NewAmount("99.05")
				currency, _ := domain.NewCurrency("USD")

				exchange.
					EXPECT().
					Convert(gomock.Any(), gomock.Eq(amount), gomock.Eq(currency)).
					Return(domain.Amount{
						Value: "470.49",
					}, nil).
					Times(1)

				return exchange
			},
			input: application.UsecaseCreatePaymentInput{
				PartnerID:        "10",
				Amount:           "99.05",
				ConsumerName:     "Oliver Tsubasa",
				ConsumerDocument: "30243434597",
			},
			paymentExpected: domain.Payment{
				ID:        "01HAW44PR1XK7B027RSFE8SAAY",
				PartnerID: "10",
				Consumer: domain.Consumer{
					Name:       "Oliver Tsubasa",
					NationalID: "30243434597",
				},
				Amount: domain.Amount{
					Value: "99.05",
				},
				ForeignAmount: domain.Amount{
					Value: "470.49",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()
			ctrl := gomock.NewController(t)

			usecase := application.NewUsecaseCreatePayment(
				test.partnerRepository(ctrl),
				test.paymentRepository(ctrl),
				test.exchange(ctrl),
			)
			paymentGot, errGot := usecase.Execute(ctx, test.input)

			assert.Equal(t, test.paymentExpected, paymentGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}