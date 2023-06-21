package cpf_test

import (
	"testing"

	"github.com/bruandreo/cpfvalidator/cpf"
	"github.com/stretchr/testify/assert"
)

func TestNewCPFWithoutSpecialCharacters(t *testing.T) {
	expectedCPF := "89613047042"

	cpf, error := cpf.NewCPF("89613047042")

	assert.Equal(t, nil, error)
	assert.Equal(t, cpf.Value, expectedCPF)
}

func TestNewCPFWithSpecialCharacters(t *testing.T) {
	expectedCPF := "896.130.470-42"

	cpf, err := cpf.NewCPF("896.130.470-42")

	assert.Equal(t, nil, err)
	assert.Equal(t, cpf.Value, expectedCPF)
}

func TestCPFCannotBeNull(t *testing.T) {
	expectError := "invalid CPF"

	_, err := cpf.NewCPF("")

	assert.EqualError(t, err, expectError)
}

func TestCPFHasInvalidLength(t *testing.T) {
	expectedError := "invalid CPF size"

	_, err := cpf.NewCPF("896.3858")

	assert.EqualError(t, err, expectedError)
}

func TestCPFCannotWithSameDigit(t *testing.T) {
	expectedError := "invalid CPF"

	_, err := cpf.NewCPF("111.111.111-11")

	assert.EqualError(t, err, expectedError)
}
