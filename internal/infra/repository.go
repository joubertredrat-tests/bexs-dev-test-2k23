package infra

import (
	"context"
	"database/sql"
	"joubertredrat/bexs-dev-test-2k23/internal/domain"
	"time"
)

type PartnerRepositoryMysql struct {
	db *sql.DB
}

func NewPartnerRepositoryMysql(db *sql.DB) PartnerRepositoryMysql {
	return PartnerRepositoryMysql{
		db: db,
	}
}

func (r PartnerRepositoryMysql) GetByID(ctx context.Context, ID string) (domain.Partner, error) {
	return domain.Partner{}, nil
}

func (r PartnerRepositoryMysql) GetByDocument(ctx context.Context, document string) (domain.Partner, error) {
	return domain.Partner{}, nil
}

func (r PartnerRepositoryMysql) Create(ctx context.Context, partner domain.Partner) (domain.Partner, error) {
	return domain.Partner{}, nil
}

type PaymentRepositoryMysql struct {
	db *sql.DB
}

func NewPaymentRepositoryMysql(db *sql.DB) PaymentRepositoryMysql {
	return PaymentRepositoryMysql{
		db: db,
	}
}

func (r PaymentRepositoryMysql) GetByID(ctx context.Context, ID string) (domain.Payment, error) {
	return domain.Payment{}, nil
}

func (r PaymentRepositoryMysql) GetDuplicated(ctx context.Context, payment domain.Payment, seconds time.Time) (domain.Payment, error) {
	return domain.Payment{}, nil
}

func (r PaymentRepositoryMysql) List(ctx context.Context, paginaton domain.Pagination) ([]domain.Payment, error) {
	return []domain.Payment{}, nil
}

func (r PaymentRepositoryMysql) Create(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	return domain.Payment{}, nil
}
