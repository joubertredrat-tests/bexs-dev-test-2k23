package application

import (
	"context"
	"joubertredrat/bexs-dev-test-2k23/internal/domain"
)

type UsecaseCreatePartner struct {
	partnerRepository domain.PartnerRepository
}

func NewUsecaseCreatePartner(r domain.PartnerRepository) UsecaseCreatePartner {
	return UsecaseCreatePartner{
		partnerRepository: r,
	}
}

func (u UsecaseCreatePartner) Execute(ctx context.Context, input UsecaseCreatePartnerInput) (domain.Partner, error) {
	currency, err := domain.NewCurrency(input.Currency)
	if err != nil {
		return domain.Partner{}, err
	}

	partnerGot, err := u.partnerRepository.GetByID(ctx, input.ID)
	if err != nil {
		return domain.Partner{}, err
	}
	if partnerGot.ID == input.ID {
		return domain.Partner{}, NewErrPartnerAlreadyExists("ID", input.ID)
	}

	partnerGot, err = u.partnerRepository.GetByDocument(ctx, input.Document)
	if err != nil {
		return domain.Partner{}, err
	}
	if partnerGot.Document == input.Document {
		return domain.Partner{}, NewErrPartnerAlreadyExists("document", input.Document)
	}

	partner := domain.NewPartner(input.ID, input.TradingName, input.Document, currency)
	return u.partnerRepository.Create(ctx, partner)
}
