package cpf

import "errors"

type CPF struct {
	Value string
}

func NewCPF(value string) (cpf *CPF, err error) {
	if isNull(value) {
		return cpf, errors.New("invalid CPF")
	}

	return &CPF{Value: value}, nil
}

func isNull(value string) bool {
	return value == ""
}
