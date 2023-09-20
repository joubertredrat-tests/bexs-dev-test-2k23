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
