package application

import (
	"context"
	"joubertredrat/bexs-dev-test-2k23/internal/domain"
	"time"
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
	duplicatedSeconds uint
}

func NewUsecaseCreatePayment(
	ptr domain.PartnerRepository,
	pyr domain.PaymentRepository,
	ex domain.Exchange,
	ds uint,
) UsecaseCreatePayment {
	return UsecaseCreatePayment{
		partnerRepository: ptr,
		paymentRepository: pyr,
		exchange:          ex,
		duplicatedSeconds: ds,
	}
}

func (u UsecaseCreatePayment) Execute(ctx context.Context, input UsecaseCreatePaymentInput) (domain.Payment, error) {
	consumer, err := domain.NewConsumer(input.ConsumerName, input.ConsumerNationalID)
	if err != nil {
		return domain.Payment{}, err
	}

	amount, err := domain.NewAmount(input.Amount)
	if err != nil {
		return domain.Payment{}, err
	}

	partnerGot, err := u.partnerRepository.GetByID(ctx, input.PartnerID)
	if err != nil {
		return domain.Payment{}, err
	}
	if partnerGot.ID != input.PartnerID {
		return domain.Payment{}, NewErrPartnerNotFound(input.PartnerID)
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

	now := time.Now()
	seconds := now.Add(-time.Second * time.Duration(u.duplicatedSeconds))

	paymentDuplcated, err := u.paymentRepository.GetDuplicated(ctx, payment, seconds)
	if err != nil {
		return domain.Payment{}, err
	}
	if paymentDuplcated.ID != "" {
		return domain.Payment{}, NewErrPaymentDuplicated(payment.PartnerID, payment.Consumer.NationalID, payment.Amount.Value)
	}

	return u.paymentRepository.Create(ctx, payment)
}

type UsecaseGetPayment struct {
	paymentRepository domain.PaymentRepository
}

func NewUsecaseGetPayment(r domain.PaymentRepository) UsecaseGetPayment {
	return UsecaseGetPayment{
		paymentRepository: r,
	}
}

func (u UsecaseGetPayment) Execute(ctx context.Context, ID string) (domain.Payment, error) {
	paymentGot, err := u.paymentRepository.GetByID(ctx, ID)
	if err != nil {
		return domain.Payment{}, err
	}
	if paymentGot.ID != ID {
		return domain.Payment{}, NewErrPaymentNotFound(ID)
	}

	return paymentGot, nil
}

type UsecaseListPayments struct {
	paymentRepository domain.PaymentRepository
}

func NewUsecaseListPayments(r domain.PaymentRepository) UsecaseListPayments {
	return UsecaseListPayments{
		paymentRepository: r,
	}
}

func (u UsecaseListPayments) Execute(ctx context.Context, pagination domain.Pagination) ([]domain.Payment, error) {
	return u.paymentRepository.List(ctx, pagination)
}
