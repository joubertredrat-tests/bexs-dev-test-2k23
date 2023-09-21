package application

import "fmt"

type ErrPartnerAlreadyExists struct {
	field string
	value string
}

func NewErrPartnerAlreadyExists(field, value string) ErrPartnerAlreadyExists {
	return ErrPartnerAlreadyExists{
		field: field,
		value: value,
	}
}

func (e ErrPartnerAlreadyExists) Error() string {
	return fmt.Sprintf("Partner already exists by [ %s ] with [ %s ]", e.field, e.value)
}

type ErrPartnerNotFound struct {
	id string
}

func NewErrPartnerNotFound(id string) ErrPartnerNotFound {
	return ErrPartnerNotFound{
		id: id,
	}
}

func (e ErrPartnerNotFound) Error() string {
	return fmt.Sprintf("Partner not found by ID [ %s ]", e.id)
}

type ErrPaymentNotFound struct {
	id string
}

func NewErrPaymentNotFound(id string) ErrPaymentNotFound {
	return ErrPaymentNotFound{
		id: id,
	}
}

func (e ErrPaymentNotFound) Error() string {
	return fmt.Sprintf("Payment not found by ID [ %s ]", e.id)
}
