package application_test

import (
	"joubertredrat/bexs-dev-test-2k23/internal/application"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrPartnerAlreadyExists(t *testing.T) {
	errExpected := "Partner already exists by [ id ] with [ 10 ]"
	errGot := application.NewErrPartnerAlreadyExists("id", "10")

	assert.Equal(t, errExpected, errGot.Error())
}

func TestErrPartnerNotFound(t *testing.T) {
	errExpected := "Partner not found by ID [ 10 ]"
	errGot := application.NewErrPartnerNotFound("10")

	assert.Equal(t, errExpected, errGot.Error())
}

func TestErrPaymentNotFound(t *testing.T) {
	errExpected := "Payment not found by ID [ 01HAW44PR1XK7B027RSFE8SAAY ]"
	errGot := application.NewErrPaymentNotFound("01HAW44PR1XK7B027RSFE8SAAY")

	assert.Equal(t, errExpected, errGot.Error())
}

func TestErrPaymentDuplicated(t *testing.T) {
	errExpected := "Payment duplicated for partner ID [ 10 ] consumer national ID [ 30243434597 ] and amount [ 99.05 ]"
	errGot := application.NewErrPaymentDuplicated("10", "30243434597", "99.05")

	assert.Equal(t, errExpected, errGot.Error())
}
