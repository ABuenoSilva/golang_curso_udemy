package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) { //Tem que iniciar com Test
	//t.Parallel() //Identifica que pode ser executado em paralelo
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"RUA Eça de Queiroz", "Rua"},
		{"Avenida Novo Horizonte", "Avenida"},
		{"Estrada dos Alvarengas", "Estrada"},
		{"Rodovia dos Bandeirantes", "Rodovia"},
		{"Praça Lauro Gomes", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if cenario.retornoEsperado != tipoDeEnderecoRecebido {
			t.Errorf("O tipo recebido é diferente do esperado!\n Esperava %s e recebeu %s", cenario.retornoEsperado, tipoDeEnderecoRecebido)
		}
	}

}
