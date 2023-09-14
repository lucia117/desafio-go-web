package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"

	"github.com/gin-gonic/gin"
)

type Router interface {
	//Metodo donde se van a definir las rutas
	MapRoutes()
}

type router struct {
	server *gin.Engine
	list   []domain.Ticket
}

func NewRouter(server *gin.Engine, list []domain.Ticket) Router {
	return &router{
		server: server,
		list:   list,
	}
}
func (router *router) MapRoutes() {
	//Instancio el repositorio
	repositorie := tickets.NewRepository(router.list)
	service := tickets.NewService(repositorie)
	handler := handler.NewHandler(service)

	//Groups
	ticketsRoutes := router.server.Group("/ticket")

	ticketsRoutes.GET("/getByCountry/:dest", handler.GetTicketsByCountry())
	ticketsRoutes.GET("/getAverage/:dest", handler.AverageDestination())

}
