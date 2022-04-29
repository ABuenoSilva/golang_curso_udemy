package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:               "/criar-usuario",
		Metodo:            http.MethodGet,
		Funcao:            controllers.CarregarTelaDeCadastroDeUsuario,
		RequerAutorizacao: false,
	},
	{
		URI:               "/usuarios",
		Metodo:            http.MethodPost,
		Funcao:            controllers.CriarUsuario,
		RequerAutorizacao: false,
	},
	{
		URI:               "/buscar-usuarios",
		Metodo:            http.MethodGet,
		Funcao:            controllers.CarregarPaginaDeUsuarios,
		RequerAutorizacao: true,
	},
	{
		URI:               "/usuarios/{usuarioID}",
		Metodo:            http.MethodGet,
		Funcao:            controllers.CarregarPerfilDoUsuario,
		RequerAutorizacao: true,
	},
	{
		URI:               "/usuarios/{usuarioID}/seguir",
		Metodo:            http.MethodPost,
		Funcao:            controllers.SeguirUsuario,
		RequerAutorizacao: true,
	},
	{
		URI:               "/usuarios/{usuarioID}/parar-de-seguir",
		Metodo:            http.MethodPost,
		Funcao:            controllers.PararDeSeguirUsuario,
		RequerAutorizacao: true,
	},
}
