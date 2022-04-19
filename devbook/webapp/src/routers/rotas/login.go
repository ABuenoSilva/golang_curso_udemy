package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasLogin = []Rota{
	{
		URI:               "/",
		Metodo:            http.MethodGet,
		Funcao:            controllers.CarregarTelaDeLogin,
		RequerAutorizacao: false,
	},
	{
		URI:               "/login",
		Metodo:            http.MethodGet,
		Funcao:            controllers.CarregarTelaDeLogin,
		RequerAutorizacao: false,
	},
	{
		URI:               "/login",
		Metodo:            http.MethodPost,
		Funcao:            controllers.FazerLogin,
		RequerAutorizacao: false,
	},
}
