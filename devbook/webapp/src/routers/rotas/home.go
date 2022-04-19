package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotaPaginaPrincial = Rota{

	URI:               "/home",
	Metodo:            http.MethodGet,
	Funcao:            controllers.CarregarPaginaPrincipal,
	RequerAutorizacao: true,
}
