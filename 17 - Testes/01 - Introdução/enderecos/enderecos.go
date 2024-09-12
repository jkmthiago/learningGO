package enderecos

import (
	"strings"
)

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia", "logradouro", "cep", "quadra", "fl", "folha", "lote"}

	enderecoLower := strings.ToLower(endereco)
	primPalavraEnd := strings.Split(enderecoLower, " ")[0]

	enderecoIsValid := false

	for _, tipo := range tiposValidos {
		if tipo == primPalavraEnd {
			enderecoIsValid = true
		}
	}

	if enderecoIsValid {
		return strings.Title(primPalavraEnd)
	}

	return "Tipo Invalido"
}
