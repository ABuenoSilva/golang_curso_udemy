package routers

import (
	"webapp/src/routers/rotas"

	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	return rotas.Configurar(mux.NewRouter())
}
