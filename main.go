package main

import (
	"net/http"

	orderapi "github.com/mrScorpio/dz/4-order-api"
)

func main() {

	srv := http.Server{
		Addr:    ":8088",
		Handler: orderapi.ProdServ(orderapi.StatusWork),
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
	//orderapi.CrTbl()
	//validationApi.CreateServer()
	// randomApi.CrWebServ()
	// concurrency.CreateRoutines()
}
