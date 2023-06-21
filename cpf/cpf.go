package cpf

import (
	"errors"
	"regexp"
	"strings"
)

type CPF struct {
	Value string
}

const (
	CPF_SIZE = 11
)

func NewCPF(value string) (cpf *CPF, err error) {
	if isNull(value) {
		return cpf, errors.New("invalid CPF")
	}

	valueClean := clear(value)

	if !hasCorrectSize(valueClean) {
		return cpf, errors.New("invalid CPF size")
	}

	if allSameDigit(valueClean) {
		return cpf, errors.New("invalid CPF")
	}

	return &CPF{Value: value}, nil
}

func isNull(value string) bool {
	return value == ""
}

func hasCorrectSize(value string) bool {
	return len(value) == CPF_SIZE
}

func clear(value string) string {
	sampleRegex := regexp.MustCompile(`\D+`)
	result := sampleRegex.ReplaceAllString(value, "")
	return result
}

func allSameDigit(value string) bool {
	digits := strings.Split(value, "")
	for _, digit := range digits {
		if digit != digits[0] {
			return false
		}
	}
	return true
}
