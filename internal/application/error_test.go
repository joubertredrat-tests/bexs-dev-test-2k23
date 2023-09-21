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
