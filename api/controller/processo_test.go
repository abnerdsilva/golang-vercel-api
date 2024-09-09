package controller_test

import (
	"github.com/stretchr/testify/require"
	"golang-vercel-api/api/controller"
	"golang-vercel-api/api/model"
	"golang-vercel-api/api/model/repository"
	"testing"
	"time"
)

func TestProcessoController_ValidProcesso(t *testing.T) {
	dbFile := "../../../DB/processos.json"

	repo := repository.NewProcessoRepository(dbFile)
	control := controller.NewProcessoController(repo)

	var processo model.Processo
	require.Error(t, control.ValidProcesso(&processo))

	processo.NrCNJ = "12345678932165498714"
	processo.DataInicio = time.Now().String()
	processo.Tribunal = model.ProcessoTribunal{
		Nome: "TJSP",
	}
	require.Nil(t, control.ValidProcesso(&processo))

	processo.NrCNJ = "1234567893216549871e"
	require.Error(t, control.ValidProcesso(&processo))

	processo.NrCNJ = "1234567"
	require.Error(t, control.ValidProcesso(&processo))

	processo.NrCNJ = ""
	require.Error(t, control.ValidProcesso(&processo))

	processo.DataInicio = ""
	require.Error(t, control.ValidProcesso(&processo))

	processo.Descricao = ""
	require.Error(t, control.ValidProcesso(&processo))

	processo.Tribunal.Nome = ""
	require.Error(t, control.ValidProcesso(&processo))
}
