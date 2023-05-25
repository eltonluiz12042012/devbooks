package rotas

import "net/http"

var rotasUsuarios = []Rota{
	{
		Uri:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: func(http.ResponseWriter, *http.Request) {

		},
		RequerAutenticacao: false,
	},

	{
		Uri:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: func(http.ResponseWriter, *http.Request) {

		},
		RequerAutenticacao: false,
	},
	{
		Uri:    "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: func(http.ResponseWriter, *http.Request) {

		},
		RequerAutenticacao: false,
	},

	{
		Uri:    "/usuarios/{usuarioId}",
		Metodo: http.MethodPut,
		Funcao: func(http.ResponseWriter, *http.Request) {

		},
		RequerAutenticacao: false,
	},

	{
		Uri:    "/usuarios/{usuarioId}",
		Metodo: http.MethodDelete,
		Funcao: func(http.ResponseWriter, *http.Request) {

		},
		RequerAutenticacao: false,
	},
}
