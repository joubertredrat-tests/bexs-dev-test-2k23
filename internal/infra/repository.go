package infra

import (
	"context"
	"database/sql"
	"joubertredrat/bexs-dev-test-2k23/internal/domain"
	"time"

	"github.com/oklog/ulid/v2"
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
	stmt, err := r.db.PrepareContext(
		ctx,
		`SELECT
			id,
			trading_name,
			document,
			currency
		FROM partners
		WHERE id = ?`,
	)
	if err != nil {
		return domain.Partner{}, err
	}
	defer stmt.Close()

	var p domain.Partner
	row := stmt.QueryRowContext(ctx, ID)
	errs := row.Scan(
		&p.ID,
		&p.TradingName,
		&p.Document,
		&p.Currency.Value,
	)
	if errs != nil {
		switch {
		case errs == sql.ErrNoRows:
			return domain.Partner{}, nil
		}
		return domain.Partner{}, err
	}

	return p, nil
}

func (r PartnerRepositoryMysql) GetByDocument(ctx context.Context, document string) (domain.Partner, error) {
	stmt, err := r.db.PrepareContext(
		ctx,
		`SELECT
			id,
			trading_name,
			document,
			currency
		FROM partners
		WHERE document = ?`,
	)
	if err != nil {
		return domain.Partner{}, err
	}
	defer stmt.Close()

	var p domain.Partner
	row := stmt.QueryRowContext(ctx, document)
	errs := row.Scan(
		&p.ID,
		&p.TradingName,
		&p.Document,
		&p.Currency.Value,
	)
	if errs != nil {
		switch {
		case errs == sql.ErrNoRows:
			return domain.Partner{}, nil
		}
		return domain.Partner{}, err
	}

	return p, nil
}

func (r PartnerRepositoryMysql) Create(ctx context.Context, partner domain.Partner) (domain.Partner, error) {
	_, err := r.db.ExecContext(
		context.Background(),
		`INSERT INTO partners (
			id,
			trading_name,
			document,
			currency
		) VALUES (?, ?, ?, ?)`,
		partner.ID,
		partner.TradingName,
		partner.Document,
		partner.Currency.Value,
	)
	if err != nil {
		return domain.Partner{}, err
	}

	return partner, nil
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
	stmt, err := r.db.PrepareContext(
		ctx,
		`SELECT
			id,
			partner_id,
			amount,
			foreign_amount,
			consumer_name,
			consumer_national_id,
			created_at
		FROM payments
		WHERE id = ?`,
	)
	if err != nil {
		return domain.Payment{}, err
	}
	defer stmt.Close()

	var p domain.Payment
	row := stmt.QueryRowContext(ctx, ID)
	errs := row.Scan(
		&p.ID,
		&p.PartnerID,
		&p.Amount.Value,
		&p.ForeignAmount.Value,
		&p.Consumer.Name,
		&p.Consumer.NationalID,
		&p.Created,
	)
	if errs != nil {
		switch {
		case errs == sql.ErrNoRows:
			return domain.Payment{}, nil
		}
		return domain.Payment{}, err
	}

	return p, nil
}

func (r PaymentRepositoryMysql) GetDuplicated(ctx context.Context, payment domain.Payment, seconds time.Time) (domain.Payment, error) {
	stmt, err := r.db.PrepareContext(
		ctx,
		`SELECT
			id,
			partner_id,
			amount,
			foreign_amount,
			consumer_name,
			consumer_national_id,
			created_at
		FROM payments
		WHERE partner_id = ? AND consumer_national_id = ? AND amount LIKE ? AND created_at >= ?`,
	)
	if err != nil {
		return domain.Payment{}, err
	}
	defer stmt.Close()

	var p domain.Payment
	row := stmt.QueryRowContext(
		ctx,
		payment.PartnerID,
		payment.Consumer.NationalID,
		payment.Amount.Value,
		DatetimeCanonical(&seconds),
	)
	errs := row.Scan(
		&p.ID,
		&p.PartnerID,
		&p.Amount.Value,
		&p.ForeignAmount.Value,
		&p.Consumer.Name,
		&p.Consumer.NationalID,
		&p.Created,
	)
	if errs != nil {
		switch {
		case errs == sql.ErrNoRows:
			return domain.Payment{}, nil
		}
		return domain.Payment{}, err
	}

	return p, nil
}

func (r PaymentRepositoryMysql) List(ctx context.Context, paginaton domain.Pagination) ([]domain.Payment, error) {
	return []domain.Payment{}, nil
}

func (r PaymentRepositoryMysql) Create(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	payment.ID = ulid.Make().String()
	payment.Created = time.Now()

	_, err := r.db.ExecContext(
		context.Background(),
		`INSERT INTO payments (
			id,
			partner_id,
			amount,
			foreign_amount,
			consumer_name,
			consumer_national_id,
			created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		payment.ID,
		payment.PartnerID,
		payment.Amount.Value,
		payment.ForeignAmount.Value,
		payment.Consumer.Name,
		payment.Consumer.NationalID,
		DatetimeCanonical(&payment.Created),
	)
	if err != nil {
		return domain.Payment{}, err
	}

	return payment, nil
}
