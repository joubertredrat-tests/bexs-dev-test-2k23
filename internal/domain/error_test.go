package domain_test

import (
	"fmt"
	"joubertredrat/bexs-dev-test-2k23/internal/domain"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrInvalidCurrency(t *testing.T) {
	currency := "BRL"
	currencies := []string{"USD", "EUR"}
	errExpected := fmt.Sprintf("Invalid currency got [ %s ], expected [ %s ]", currency, strings.Join(currencies, ", "))
	errGot := domain.NewErrInvalidCurrency(currency, currencies)

	assert.Equal(t, errExpected, errGot.Error())
}
