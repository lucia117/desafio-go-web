package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/bootcamp-go/desafio-go-web/cmd/server/router"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/gin-gonic/gin"
)

func main() {

	// Cargo csv desde el archivo especificado.
	list, err := LoadTicketsFromFile("../../tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

	server := gin.Default()

	// Configura el enrutador y mapea las rutas.
	router := router.NewRouter(server, list)
	router.MapRoutes()

	server.GET("/ping", func(ctx *gin.Context) { ctx.String(200, "pong") })

	// Inicia el servidor Gin.
	if err := server.Run(); err != nil {
		panic(err)
	}

}

// LoadTicketsFromFile carga los datos de los tickets desde un archivo CSV
// y devuelve una lista de tickets y un error si ocurriera.
func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
