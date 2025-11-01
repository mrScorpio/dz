package validationApi

import (
	"log"
	"net/http"
)

func CreateServer() {
	conf := LoadConfig()
	mux := http.NewServeMux()
	NewValHandler(mux, ValHandlerDeps{
		Config: conf,
	})

	srv := http.Server{
		Addr:    ":8086",
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
	}
}
