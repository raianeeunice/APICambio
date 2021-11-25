package modelos

import (
	"time"
)

type Deposito struct {
	ID            uint64    `json:"id,omitempty"`
	ValorDeposito float64   `json:"valorDeposito,omitempty"`
	FeitoEm       *time.Time `json:"feitoEm,omitempty"`
}
