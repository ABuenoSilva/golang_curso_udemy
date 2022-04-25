package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPublicacoes = []Rota{
	{
		URI:               "/publicacoes",
		Metodo:            http.MethodPost,
		Funcao:            controllers.CriarPublicacao,
		RequerAutorizacao: true,
	},
	{
		URI:               "/publicacoes/{publicacaoID}/curtir",
		Metodo:            http.MethodPost,
		Funcao:            controllers.CurtirPublicacao,
		RequerAutorizacao: true,
	},
}
