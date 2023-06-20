package cpf

type CPF struct {
	Value string
}

func NewCPF(value string) (*CPF, error) {
	return &CPF{Value: value}, nil
}
