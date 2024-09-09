package model

import "github.com/google/uuid"

type Processo struct {
	UID           uuid.UUID            `json:"uid"`
	NrCNJ         string               `json:"nr_cnj"`
	DataInicio    string               `json:"data_inicio"`
	Descricao     string               `json:"descricao"`
	Tribunal      ProcessoTribunal     `json:"tribunal_origem"`
	Status        string               `json:"status_processo"`
	NrInstancia   int                  `json:"nr_instancia"`
	VrCausa       float64              `json:"vr_causa"`
	Envolvidos    []ProcessoEnvolvidos `json:"envolvidos"`
	Movimentacoes []ProcessoHistorico  `json:"movimentacoes"`
}

type ProcessoEnvolvidos struct {
	Nome string `json:"nome"`
	Tipo int    `json:"tipo"`
}

type ProcessoTribunal struct {
	Nome  string `json:"nome"`
	Tipo  string `json:"tipo"`
	Local string `json:"local"`
}

type ProcessoHistorico struct {
	UID              uuid.UUID `json:"uid"`
	DataMovimentacao string    `json:"data_movimentacao"`
	Descricao        string    `json:"descricao"`
}
