package main

import (
	"context"
	"listaTarefas/data"
	"listaTarefas/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

// go mod init listaTarefas
// go get github.com/jackc/pgx/v5 (para baixar o driver do banco)
func main() {
	cnx, err := data.ConectarBanco()
	if err != nil {
		panic(err)
	}
	defer cnx.Close(context.Background())
	err = data.CriaTabela(cnx)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	tarefasHandler := &handler.TarefasHandler{
		Db: cnx,
	}
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/tarefas", tarefasHandler.HandlerListarTarefas)
	r.POST("/tarefas", tarefasHandler.HandlerIncluirTarefa)
	r.PUT("/tarefas/:id", tarefasHandler.HandlerEditarTarefa)
	r.DELETE("/tarefas/:id", tarefasHandler.HandlerExcluirTarefa)
	r.PUT("/tarefas/:id/subir", tarefasHandler.HandlerSubirTarefa)
	r.PUT("/tarefas/:id/descer", tarefasHandler.HandlerDescerTarefa)
	r.Run()
}
