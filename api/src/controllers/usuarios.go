package controllers

import "net/http"

//CriarUsuario cria um novo usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!"))
}

//BuscarUsuarios Busca todos usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("BUscando todos usuários!"))
}

//BuscarUsuario Busca Todos Usuarios
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário!"))
}

//AtualizarUsuario atualiza um usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário!"))
}

//DeletarUsuario Deleta um usuario
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
