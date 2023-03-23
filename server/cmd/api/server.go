package main

import (
	"net/http"

	"github.com/luiscovelo/goexpert-client-server-api/server/controller"
	"github.com/luiscovelo/goexpert-client-server-api/server/database"
	"github.com/luiscovelo/goexpert-client-server-api/server/repository"
	"github.com/luiscovelo/goexpert-client-server-api/server/service"
)

func main() {
	db, err := database.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.NewQuotationRepository(db.DB)
	service := service.NewQuotationService(repo)
	controller := controller.NewQuotationController(service)

	server := http.NewServeMux()
	server.HandleFunc("/cotacao", controller.GetQuotation)

	err = http.ListenAndServe(":8080", server)
	if err != nil {
		panic(err)
	}
}
