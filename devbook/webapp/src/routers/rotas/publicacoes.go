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
	{
		URI:               "/publicacoes/{publicacaoID}/atualizar",
		Metodo:            http.MethodGet,
		Funcao:            controllers.CarregarPaginaDeEdicaoDePublicacao,
		RequerAutorizacao: true,
	},
	{
		URI:               "/publicacoes/{publicacaoID}/deletar",
		Metodo:            http.MethodDelete,
		Funcao:            controllers.CurtirPublicacao,
		RequerAutorizacao: true,
	},
	{
		URI:               "/publicacoes/{publicacaoID}",
		Metodo:            http.MethodPut,
		Funcao:            controllers.AtualizarPublicacao,
		RequerAutorizacao: true,
	},
	{
		URI:               "/publicacoes/{publicacaoID}",
		Metodo:            http.MethodDelete,
		Funcao:            controllers.DeletarPublicacao,
		RequerAutorizacao: true,
	},
}
