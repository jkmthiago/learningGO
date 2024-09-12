// Teste De Unidade
// package enderecos_test

package enderecos_test

import (
	. "testes/enderecos"
	"testing"
)

type cenarioDeTest struct {
	endereco       string
	expectedReturn string
}

func TestTipoDeEndereco(t *testing.T) {

	t.Parallel()

	cenariosDeTest := []cenarioDeTest{
		{"Rua Tomé de Souza", "Rua"},
		{"CEP 68506-660 - Quadra Dez (Fl-29), Nova Marabá, Marabá - PA", "Cep"},
		{"Quadra Dezenove (Fl.28), Nova Marabá, Marabá - PA", "Quadra"},
		{"Quadra Quatro (Fl.28), Nova Marabá, Marabá - PA", "Quadra"},
		{"CEP 68506-280 - Quadra Vinte e Oito (Fl.28), Nova Marabá, Marabá - PA", "Cep"},
		{"Fl 33, Nova Marabá, Marabá - PA", "Fl"},
	}

	for _, cenario := range cenariosDeTest {
		received := TipoDeEndereco(cenario.endereco)
		if received != cenario.expectedReturn {
			t.Errorf("O endereço %s retornou um tipo invalido ou ainda não cadastrado. '%s'\nO experado era: %s", cenario.endereco, received, cenario.expectedReturn)
		}
	}
}

// func TestTipoDeEndereco(t *testing.T)  {
// 	enderecoTeste := "Rua Tomé de Souza"
// 	expected := "Rua"
// 	received := TipoDeEndereco(enderecoTeste)

// 	if received != expected{
// 		t.Errorf("O endereço %s retornou um tipo invalido ou ainda não cadastrado. '%s'\nO experado era: %s", enderecoTeste, received, expected)
// 	}
// }

func TestTipoDeEndereco2(t *testing.T) {
	t.Parallel()
	enderecosTeste := []string{
		"Rua Tomé de Souza",
		"CEP 68506-660 - Quadra Dez (Fl-29), Nova Marabá, Marabá - PA",
		"Quadra Dezenove (Fl.28), Nova Marabá, Marabá - PA",
		"Quadra Quatro (Fl.28), Nova Marabá, Marabá - PA",
		"CEP 68506-280 - Quadra Vinte e Oito (Fl.28), Nova Marabá, Marabá - PA",
		"Fl 33, Nova Marabá, Marabá - PA",
		"Avenida 33, Nova Marabá, Marabá - PA",
		"Rodovia Transamazônica, Fl 33, Nova Marabá, Marabá - PA",
		"Estrada de Ferro Carajá, Fl 33, Nova Marabá, Marabá - PA",
		"Logradouro Boa Esperança",
		"Folha 33, Nova Marabá, Marabá - PA",
		"Lote 04, Fl 33, Nova Marabá, Marabá - PA",
		"",
	}

	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia", "Logradouro", "Cep", "Quadra", "Fl", "Folha", "Lote", "Tipo Invalido"}

	// received := TipoDeEndereco(enderecoTeste)

	received := []string{}

	for _, endereco := range enderecosTeste {
		received = append(received, TipoDeEndereco(endereco))
	}

	// if received != expected {
	// 	t.Error("O tipo de endereço fornecido é diferente do experado.")
	// }

	enderecoPosition := 0

	for _, received := range received {
		asWeExpected := true
		countOfVerified := 0
		for _, valid := range tiposValidos {
			countOfVerified++
			if received == valid {
				break
			} else if countOfVerified == len(tiposValidos) && received != valid {
				asWeExpected = false
			}
		}

		if !asWeExpected {
			t.Errorf("O endereço %s retornou um tipo invalido ou ainda não cadastrado.", enderecosTeste[enderecoPosition])
		}
		enderecoPosition++
	}
}
