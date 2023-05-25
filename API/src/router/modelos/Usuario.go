package modelos

import "time"

type Usuario struct {
	ID       uint64    `json:"id, omitempyt"`
	Nome     string    `json:"nome,omitempyt"`
	Nick     string    `json:"nick,omitempyt"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempyt"`
	CriadoEm time.Time `json:"CriadoEm,omitempyt"`
}
