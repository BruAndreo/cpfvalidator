package cpf_test

import (
	"testing"

	"github.com/bruandreo/cpfvalidator/cpf"
	"github.com/stretchr/testify/assert"
)

func TestNewCPF(t *testing.T) {
	expectedCPF := "896.130.470-42"

	cpf, error := cpf.NewCPF("896.130.470-42")

	assert.Equal(t, nil, error)
	assert.Equal(t, cpf.Value, expectedCPF)
}
