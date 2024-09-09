package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang-vercel-api/api/model"
	"golang-vercel-api/api/model/repository"
	"golang-vercel-api/api/view"
	"log"
	"net/http"
	"strconv"
)

func NewProcessoController(processoRepository repository.ProcessoRepository) ProcessoControllerInterface {
	return &processoController{processoRepository}
}

type ProcessoControllerInterface interface {
	CreateProcesso(c *gin.Context)
	GetProcessos(c *gin.Context)
	GetProcesso(c *gin.Context)
	ValidProcesso(processo *model.Processo) error
}

type processoController struct {
	processoRepository repository.ProcessoRepository
}

func (pc *processoController) CreateProcesso(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST")

	var tempProcesso model.Processo
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&tempProcesso)
	if err != nil {
		fmt.Printf("error %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	tempProcesso.UID = uuid.New()

	err = pc.ValidProcesso(&tempProcesso)
	if err != nil {
		err2 := view.ParseError(fmt.Sprintf("Erro na validação do processo -> %s", err), http.StatusBadRequest)
		c.JSONP(err2.Code, err2)
		return
	}

	tNrCNJ := tempProcesso.NrCNJ
	tempProcesso.NrCNJ = tNrCNJ[:7] + "." + tNrCNJ[7:9] + "." + tNrCNJ[9:13] + "." + tNrCNJ[13:14] + "." + tNrCNJ[14:16] + "." + tNrCNJ[16:20]

	var movimentacoes []model.ProcessoHistorico
	for _, mov := range tempProcesso.Movimentacoes {
		mov.UID = uuid.New()

		movimentacoes = append(movimentacoes, mov)
	}

	tempProcesso.Movimentacoes = movimentacoes

	createdProcess, err := pc.processoRepository.CreateProcesso(&tempProcesso)
	if err != nil {
		err2 := view.ParseError(fmt.Sprintf("Erro na criação do processo -> %s", err), http.StatusBadRequest)
		c.JSONP(err2.Code, err2)
		return
	}

	c.JSONP(http.StatusOK, view.ConvertProcessoToView(createdProcess))
}

func (pc *processoController) GetProcessos(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, OPTIONS")

	processos, err := pc.processoRepository.GetProcessos()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro consulta processos"})
		return
	}

	c.JSON(http.StatusOK, processos)
}

func (pc *processoController) GetProcesso(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, OPTIONS")

	uid := c.Param("id")

	processo, err := pc.processoRepository.GetProcesso(uid)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro consulta processo"})
		return
	}

	c.JSON(http.StatusOK, processo)
}

func (pc *processoController) ValidProcesso(processo *model.Processo) error {
	if processo.NrCNJ == "" {
		return fmt.Errorf("nr_cnj é um campo obrigatório")
	}
	if len(processo.NrCNJ) != 20 {
		return fmt.Errorf("nr_cnj deve ter 20 caracteres")
	}

	num := []byte(processo.NrCNJ)
	for i := 0; i < len(num); i++ {
		tNum := fmt.Sprintf("%v", num[i])
		tNumParsed, _ := strconv.Atoi(tNum)
		if tNumParsed < 48 || tNumParsed > 57 {
			return fmt.Errorf("nr_cnj deve ter apenas números")
		}
	}

	if processo.DataInicio == "" {
		return fmt.Errorf("data_inicio é um campo obrigatório")
	}
	if processo.Tribunal.Nome == "" {
		return fmt.Errorf("tribunal é um campo obrigatório")
	}

	return nil
}
