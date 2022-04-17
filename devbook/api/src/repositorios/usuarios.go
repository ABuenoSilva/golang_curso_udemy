package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"errors"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {

	statement, erro := repositorio.db.Prepare("insert into usuarios(nome, nick, email, senha) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?",
		nomeOuNick,
		nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

func (repositorio usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?", ID)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	} else {
		return modelos.Usuario{}, errors.New("o id enviado não foi encontrado no banco de dados")
	}

	return usuario, nil
}

func (repositorio usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare("update usuarios set nome = ?, nick = ?, email=? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if linhasAfetadas, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	} else if qtdLinhas, erro := linhasAfetadas.RowsAffected(); erro != nil {
		return erro
	} else if qtdLinhas < 1 {
		return errors.New("não foi encontrado o ID de usuário para alteração ou os dados enviados são iguais aos atuais")
	}

	return nil

}

func (repositorio usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if linhasAfetadas, erro := statement.Exec(ID); erro != nil {
		return erro
	} else if qtdLinhas, erro := linhasAfetadas.RowsAffected(); erro != nil {
		return erro
	} else if qtdLinhas < 1 {
		return errors.New("não foi encontrado o ID de usuário para exclusão")
	}

	return nil
}

func (repositorio usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Senha,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	} else {
		return modelos.Usuario{}, errors.New("o id enviado não foi encontrado no banco de dados")
	}

	return usuario, nil
}

func (repositorio usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from seguidores where usuario_id =  ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query(
		`select id, u.nome, u.nick, u.email, u.criadoEm 
		 from usuarios u inner join seguidores s 
		 on u.id = s.seguidor_id and s.usuario_id = ?`,
		usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var seguidores []modelos.Usuario
	for linhas.Next() {
		var seguidor modelos.Usuario
		if erro = linhas.Scan(
			&seguidor.ID,
			&seguidor.Nome,
			&seguidor.Nick,
			&seguidor.Email,
			&seguidor.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		seguidores = append(seguidores, seguidor)
	}

	return seguidores, nil

}

func (repositorio usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query(
		"select id, u.nome, u.nick, u.email, u.criadoEm from usuarios u inner join seguidores s on u.id = s.usuario_id and s.seguidor_id = ?", usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var seguidores []modelos.Usuario
	for linhas.Next() {
		var seguidor modelos.Usuario
		if erro = linhas.Scan(
			&seguidor.ID,
			&seguidor.Nome,
			&seguidor.Nick,
			&seguidor.Email,
			&seguidor.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		seguidores = append(seguidores, seguidor)
	}

	return seguidores, nil

}

func (repositorio usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linhas, erro := repositorio.db.Query("select senha from usuarios where id = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.Senha,
		); erro != nil {
			return "", erro
		}
	} else {
		return "", errors.New("o id enviado não foi encontrado no banco de dados")
	}

	return usuario.Senha, nil
}

func (repositorio usuarios) AtualizarSenha(usuarioID uint64, novaSenha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(novaSenha, usuarioID); erro != nil {
		return erro
	}

	return nil
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}
