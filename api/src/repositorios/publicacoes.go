package repositorios

import (
	"api/api/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositorio de publicacoes
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um novo repositorio de publicacoes
func NovoRepositorioDePublicoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO PUBLICACOES (TITULO, CONTEUDO, AUTOR_ID) VALUES (?,?,?)")
	if erro != nil {
		return 0, erro
	}

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

// BuscarPorID traz uma íunica publicacão do banco de dados
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(
		`SELECT P.*, U.NICK 
		FROM PUBLICACOES P INNER JOIN USUARIOS U 
		ON U.ID = P.AUTOR_ID 
		WHERE P.ID = ?`, publicacaoID)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil

}

// Buscar retorna todas as publicaoes do feed do usuário
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(
		`SELECT DISTINCT P.*, U.NICK 
		FROM PUBLICACOES P INNER JOIN USUARIOS U 
		ON U.ID = P.AUTOR_ID 
		INNER JOIN SEGUIDORES S
		ON P.AUTOR_ID = S.USUARIO_ID 
		WHERE U.ID = ?
		OR S.SEGUIDOR_ID = ?`, usuarioID, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil

}

// Atualizar altera os dados de uma publicação no banco de dados
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("UPDATE PUBLICACOES SET TITULO = ?, CONTEUDO = ?  WHERE ID = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil

}

func (repositorio Publicacoes) Deletar(publicacaoID uint64) error{
	statement, erro := repositorio.db.Prepare("DELETE FROM PUBLICACOES WHERE ID = ?",)
	if erro != nil{
		return erro		
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil{
		return erro
	}

	return nil

}

//BuscarPorUsuario traz todas as publicacoes de um usuário específico
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error){
	linhas, erro := repositorio.db.Query(`
		SELECT P.*, U.NICK FROM PUBLICACOES P
		JOIN USUARIOS U ON U.ID = P.AUTOR_ID
		WHERE P.AUTOR_ID = ?
	`, usuarioID,
	)
	if erro != nil{
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

//Curtir adiciona uma curtida na publicacao
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error{
	statement, erro := repositorio.db.Prepare("UPDATE PUBLICACOES SET CURTIDAS = CURTIDAS + 1 WHERE ID = ?")

	if erro != nil{
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil{
		return erro
	}

	return nil

}

//Descurtir subtrai uma curtida na publicacao
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error{
	statement, erro := repositorio.db.Prepare(`
		UPDATE PUBLICACOES SET CURTIDAS = 
			CASE WHEN CURTIDAS > 0 THEN
				CURTIDAS - 1
			ELSE CURTIDAS END 
		 WHERE ID = ?`)

	if erro != nil{
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil{
		return erro
	}

	return nil

}