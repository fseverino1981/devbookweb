package controllers

import (
	"api/api/src/autenticacao"
	"api/api/src/banco"
	"api/api/src/modelos"
	"api/api/src/repositorios"
	"api/api/src/respostas"
	"api/api/src/seguranca"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login é responsável por autenticar um usuário na api
func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	w.Write([]byte(token))
}
