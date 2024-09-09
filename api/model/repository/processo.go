package repository

import (
	"encoding/json"
	"fmt"
	"golang-vercel-api/api/model"
	"io/ioutil"
	"log"
	"os"
)

type ProcessoRepository interface {
	CreateProcesso(processo *model.Processo) (*[]model.Processo, error)
	GetProcessos() ([]model.Processo, error)
	GetProcesso(uid string) (model.Processo, error)
}

type processoRepository struct {
	DB string
}

func NewProcessoRepository(db string) ProcessoRepository {
	return &processoRepository{
		DB: db,
	}
}

func (p *processoRepository) CreateProcesso(processo *model.Processo) (*[]model.Processo, error) {
	process, err := p.GetProcessos()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, proc := range process {
		if proc.NrCNJ == processo.NrCNJ {
			return nil, fmt.Errorf("processo já cadastrado")
		}
	}

	process = append(process, *processo)

	file, _ := json.MarshalIndent(process, "", " ")

	_ = os.WriteFile(p.DB, file, 0644)
	return &process, nil
}

func (p *processoRepository) GetProcessos() ([]model.Processo, error) {
	arr := make([]model.Processo, 0)

	fs, err := ioutil.ReadFile(p.DB)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(fs) > 2 {
		var processos []model.Processo
		err = json.Unmarshal(fs, &processos)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		for _, value := range processos {
			arr = append(arr, value)
		}

	}

	return arr, nil
}

func (p *processoRepository) GetProcesso(uid string) (model.Processo, error) {
	processos, err := p.GetProcessos()
	if err != nil {
		log.Println(err)
		return model.Processo{}, err
	}

	var processo model.Processo
	for _, proc := range processos {
		if proc.UID.String() == uid {
			processo = proc
			continue
		}
	}

	return processo, nil
}
