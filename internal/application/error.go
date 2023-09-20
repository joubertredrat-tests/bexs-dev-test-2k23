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
