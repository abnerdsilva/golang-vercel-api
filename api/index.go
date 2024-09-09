package handler

import (
	"fmt"
	//"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"golang-vercel-api/api/controller"
	"golang-vercel-api/api/model/repository"
)

func main() {
	dbFile := "../DB/processos.json"

	processRepo := repository.NewProcessoRepository(dbFile)
	processControll := controller.NewProcessoController(processRepo)

	router := gin.Default()
	//router.Use(static.Serve("/", static.LocalFile("../build", false)))

	router.GET("/process", processControll.GetProcessos)
	router.GET("/process/:id", processControll.GetProcesso)
	router.POST("/process", processControll.CreateProcesso)

	err := router.Run(":8787")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
}
