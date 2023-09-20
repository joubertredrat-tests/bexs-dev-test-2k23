package domain

import (
	"fmt"
	"strings"
)

type ErrInvalidCurrency struct {
	got      string
	expected []string
}

func NewErrInvalidCurrency(got string, expected []string) ErrInvalidCurrency {
	return ErrInvalidCurrency{
		got:      got,
		expected: expected,
	}
}

func (e ErrInvalidCurrency) Error() string {
	return fmt.Sprintf("Invalid currency got [ %s ], expected [ %s ]", e.got, strings.Join(e.expected, ", "))
}
