package application_test

import (
	"context"
	"errors"
	"joubertredrat/bexs-dev-test-2k23/internal/application"
	"joubertredrat/bexs-dev-test-2k23/internal/domain"
	"joubertredrat/bexs-dev-test-2k23/pkg/mock"
	"testing"
	"time"

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
			name: "Test create partner with invalid currency",
			partnerRepository: func(ctrl *gomock.Controller) domain.PartnerRepository {
				return mock.NewMockPartnerRepository(ctrl)
			},
			input: application.UsecaseCreatePartnerInput{
				ID:          "10",
				TradingName: "International Ecommerce",
				Document:    "1284498339812/0001",
				Currency:    "BRL",
			},
			partnerExpected: domain.Partner{},
			errExpected:     domain.NewErrInvalidCurrency("BRL", domain.Currencies()),
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
		duplicatedSeconds uint
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

				repository.
					EXPECT().
					GetDuplicated(gomock.Any(), gomock.AssignableToTypeOf(domain.Payment{}), gomock.AssignableToTypeOf(time.Time{})).
					Return(domain.Payment{}, nil).
					Times(1)

				repository.
					EXPECT().
					Create(gomock.Any(), gomock.AssignableToTypeOf(domain.Payment{})).
					Return(domain.Payment{
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
			duplicatedSeconds: 120,
			input: application.UsecaseCreatePaymentInput{
				PartnerID:          "10",
				Amount:             "99.05",
				ConsumerName:       "Oliver Tsubasa",
				ConsumerNationalID: "30243434597",
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
			errExpected: nil,
		},
		{
			name: "Test create payment with invalid consumer national ID",
			partnerRepository: func(ctrl *gomock.Controller) domain.PartnerRepository {
				return mock.NewMockPartnerRepository(ctrl)
			},
			paymentRepository: func(ctrl *gomock.Controller) domain.PaymentRepository {
				return mock.NewMockPaymentRepository(ctrl)
			},
			exchange: func(ctrl *gomock.Controller) domain.Exchange {
				return mock.NewMockExchange(ctrl)
			},
			duplicatedSeconds: 120,
			input: application.UsecaseCreatePaymentInput{
				PartnerID:          "10",
				Amount:             "99.05",
				ConsumerName:       "Oliver Tsubasa",
				ConsumerNationalID: "302434597",
			},
			paymentExpected: domain.Payment{},
			errExpected:     errors.New("Consumer national ID expect 11 digits, got [ 302434597 ]"),
		},
		{
			name: "Test create payment with invalid amount",
			partnerRepository: func(ctrl *gomock.Controller) domain.PartnerRepository {
				return mock.NewMockPartnerRepository(ctrl)
			},
			paymentRepository: func(ctrl *gomock.Controller) domain.PaymentRepository {
				return mock.NewMockPaymentRepository(ctrl)
			},
			exchange: func(ctrl *gomock.Controller) domain.Exchange {
				return mock.NewMockExchange(ctrl)
			},
			duplicatedSeconds: 120,
			input: application.UsecaseCreatePaymentInput{
				PartnerID:          "10",
				Amount:             "990.5",
				ConsumerName:       "Oliver Tsubasa",
				ConsumerNationalID: "30243434597",
			},
			paymentExpected: domain.Payment{},
			errExpected:     errors.New("Amount expect valid value, got [ 990.5 ]"),
		},
		{
			name: "Test create payment with unknown error on get by id from partner repository",
			partnerRepository: func(ctrl *gomock.Controller) domain.PartnerRepository {
				repository := mock.NewMockPartnerRepository(ctrl)

				repository.
					EXPECT().
					GetByID(gomock.Any(), gomock.Eq("10")).
					Return(domain.Partner{}, errors.New("database gone")).
					Times(1)

				return repository
			},
			paymentRepository: func(ctrl *gomock.Controller) domain.PaymentRepository {
				return mock.NewMockPaymentRepository(ctrl)
			},
			exchange: func(ctrl *gomock.Controller) domain.Exchange {
				return mock.NewMockExchange(ctrl)
			},
			duplicatedSeconds: 120,
			input: application.UsecaseCreatePaymentInput{
				PartnerID:          "10",
				Amount:             "99.05",
				ConsumerName:       "Oliver Tsubasa",
				ConsumerNationalID: "30243434597",
			},
			paymentExpected: domain.Payment{},
			errExpected:     errors.New("database gone"),
		},
		{
			name: "Test create payment with partner ID not found from repository",
			partnerRepository: func(ctrl *gomock.Controller) domain.PartnerRepository {
				repository := mock.NewMockPartnerRepository(ctrl)

				repository.
					EXPECT().
					GetByID(gomock.Any(), gomock.Eq("10")).
					Return(domain.Partner{}, nil).
					Times(1)

				return repository
			},
			paymentRepository: func(ctrl *gomock.Controller) domain.PaymentRepository {
				return mock.NewMockPaymentRepository(ctrl)
			},
			exchange: func(ctrl *gomock.Controller) domain.Exchange {
				return mock.NewMockExchange(ctrl)
			},
			duplicatedSeconds: 120,
			input: application.UsecaseCreatePaymentInput{
				PartnerID:          "10",
				Amount:             "99.05",
				ConsumerName:       "Oliver Tsubasa",
				ConsumerNationalID: "30243434597",
			},
			paymentExpected: domain.Payment{},
			errExpected:     application.NewErrPartnerNotFound("10"),
		},
		{
			name: "Test create payment with unknown error from exchange convert",
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
				return mock.NewMockPaymentRepository(ctrl)
			},
			exchange: func(ctrl *gomock.Controller) domain.Exchange {
				exchange := mock.NewMockExchange(ctrl)

				amount, _ := domain.NewAmount("99.05")
				currency, _ := domain.NewCurrency("USD")

				exchange.
					EXPECT().
					Convert(gomock.Any(), gomock.Eq(amount), gomock.Eq(currency)).
					Return(domain.Amount{}, errors.New("service down")).
					Times(1)

				return exchange
			},
			duplicatedSeconds: 120,
			input: application.UsecaseCreatePaymentInput{
				PartnerID:          "10",
				Amount:             "99.05",
				ConsumerName:       "Oliver Tsubasa",
				ConsumerNationalID: "30243434597",
			},
			paymentExpected: domain.Payment{},
			errExpected:     errors.New("service down"),
		},
		{
			name: "Test create payment with unknown error on get duplicated from payment repository",
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

				repository.
					EXPECT().
					GetDuplicated(gomock.Any(), gomock.AssignableToTypeOf(domain.Payment{}), gomock.AssignableToTypeOf(time.Time{})).
					Return(domain.Payment{}, errors.New("database gone")).
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
			duplicatedSeconds: 120,
			input: application.UsecaseCreatePaymentInput{
				PartnerID:          "10",
				Amount:             "99.05",
				ConsumerName:       "Oliver Tsubasa",
				ConsumerNationalID: "30243434597",
			},
			paymentExpected: domain.Payment{},
			errExpected:     errors.New("database gone"),
		},
		{
			name: "Test create payment with duplicated payment found",
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

				repository.
					EXPECT().
					GetDuplicated(gomock.Any(), gomock.AssignableToTypeOf(domain.Payment{}), gomock.AssignableToTypeOf(time.Time{})).
					Return(domain.Payment{
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
			duplicatedSeconds: 120,
			input: application.UsecaseCreatePaymentInput{
				PartnerID:          "10",
				Amount:             "99.05",
				ConsumerName:       "Oliver Tsubasa",
				ConsumerNationalID: "30243434597",
			},
			paymentExpected: domain.Payment{},
			errExpected:     application.NewErrPaymentDuplicated("10", "30243434597", "99.05"),
		},
		{
			name: "Test create payment with unknown error on create from payment repository",
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

				repository.
					EXPECT().
					GetDuplicated(gomock.Any(), gomock.AssignableToTypeOf(domain.Payment{}), gomock.AssignableToTypeOf(time.Time{})).
					Return(domain.Payment{}, nil).
					Times(1)

				repository.
					EXPECT().
					Create(gomock.Any(), gomock.AssignableToTypeOf(domain.Payment{})).
					Return(domain.Payment{}, errors.New("database gone")).
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
			duplicatedSeconds: 120,
			input: application.UsecaseCreatePaymentInput{
				PartnerID:          "10",
				Amount:             "99.05",
				ConsumerName:       "Oliver Tsubasa",
				ConsumerNationalID: "30243434597",
			},
			paymentExpected: domain.Payment{},
			errExpected:     errors.New("database gone"),
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
				test.duplicatedSeconds,
			)
			paymentGot, errGot := usecase.Execute(ctx, test.input)

			assert.Equal(t, test.paymentExpected, paymentGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}

func TestUsecaseGetPayment(t *testing.T) {
	tests := []struct {
		name              string
		paymentRepository func(ctrl *gomock.Controller) domain.PaymentRepository
		ID                string
		paymentExpected   domain.Payment
		errExpected       error
	}{
		{
			name: "Test get payment with success",
			paymentRepository: func(ctrl *gomock.Controller) domain.PaymentRepository {
				repository := mock.NewMockPaymentRepository(ctrl)

				repository.
					EXPECT().
					GetByID(gomock.Any(), gomock.Eq("01HAW44PR1XK7B027RSFE8SAAY")).
					Return(domain.Payment{
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
					}, nil).
					Times(1)

				return repository
			},
			ID: "01HAW44PR1XK7B027RSFE8SAAY",
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
			errExpected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()
			ctrl := gomock.NewController(t)

			usecase := application.NewUsecaseGetPayment(test.paymentRepository(ctrl))
			paymentGot, errGot := usecase.Execute(ctx, test.ID)

			assert.Equal(t, test.paymentExpected, paymentGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}

func TestUsecaseListPayments(t *testing.T) {
	tests := []struct {
		name              string
		paymentRepository func(ctrl *gomock.Controller) domain.PaymentRepository
		pagination        domain.Pagination
		paymentsExpected  []domain.Payment
		errExpected       error
	}{
		{
			name: "Test list payments with success",
			paymentRepository: func(ctrl *gomock.Controller) domain.PaymentRepository {
				repository := mock.NewMockPaymentRepository(ctrl)

				pagination := domain.NewPagination(0, 10)

				repository.
					EXPECT().
					List(gomock.Any(), gomock.Eq(pagination)).
					Return([]domain.Payment{
						domain.Payment{
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
						domain.Payment{
							ID:        "01HAW9KE9342952B9X9FAC147G",
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
					}, nil).
					Times(1)

				return repository
			},
			pagination: domain.NewPagination(0, 10),
			paymentsExpected: []domain.Payment{
				domain.Payment{
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
				domain.Payment{
					ID:        "01HAW9KE9342952B9X9FAC147G",
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
			errExpected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()
			ctrl := gomock.NewController(t)

			usecase := application.NewUsecaseListPayments(test.paymentRepository(ctrl))
			paymentsGot, errGot := usecase.Execute(ctx, test.pagination)

			assert.Equal(t, test.paymentsExpected, paymentsGot)
			assert.Equal(t, test.errExpected, errGot)
		})
	}
}
