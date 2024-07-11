package rotas


import (
	"api/api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API
=======
import "net/http"

//Rota representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}


// Configurar caloca todas as rotas dentro de um router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasDeUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}

	return r
}
