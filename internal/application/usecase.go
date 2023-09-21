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

type UsecaseCreatePayment struct {
	partnerRepository domain.PartnerRepository
	paymentRepository domain.PaymentRepository
	exchange          domain.Exchange
}

func NewUsecaseCreatePayment(
	ptr domain.PartnerRepository,
	pyr domain.PaymentRepository,
	ex domain.Exchange,
) UsecaseCreatePayment {
	return UsecaseCreatePayment{
		partnerRepository: ptr,
		paymentRepository: pyr,
		exchange:          ex,
	}
}

func (u UsecaseCreatePayment) Execute(ctx context.Context, input UsecaseCreatePaymentInput) (domain.Payment, error) {
	partnerGot, err := u.partnerRepository.GetByID(ctx, input.PartnerID)
	if err != nil {
		return domain.Payment{}, err
	}
	if partnerGot.ID != input.PartnerID {
		return domain.Payment{}, NewErrPartnerNotFound(input.PartnerID)
	}

	consumer, err := domain.NewConsumer(input.ConsumerName, input.ConsumerDocument)
	if err != nil {
		return domain.Payment{}, err
	}

	amount, err := domain.NewAmount(input.Amount)
	if err != nil {
		return domain.Payment{}, err
	}
	foreignAmount, err := u.exchange.Convert(ctx, amount, partnerGot.Currency)
	if err != nil {
		return domain.Payment{}, err
	}

	payment := domain.Payment{
		PartnerID:     input.PartnerID,
		Consumer:      consumer,
		Amount:        amount,
		ForeignAmount: foreignAmount,
	}

	return u.paymentRepository.Create(ctx, payment)
}
