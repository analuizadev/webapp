package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotaPagPrincipal = Rota{
		URI:    "/home",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarHome,
		RequerAutenticacao: true,
}