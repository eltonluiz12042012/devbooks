package modelos

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) Preparar() error {
	fmt.Println("Chamando pacote usuário")
	if erro := usuario.validar(); erro != nil {
		return erro
	}
	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O campo nome é obrigatório e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O campo nick é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O campo email é obrigatório e não pode estar em branco")
	}

	if usuario.Senha == "" {
		return errors.New("O campo senha é obrigatório e não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
