package orderapi

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func ProdServ() {
	conf := LoadConfig()
	db := NewDb(conf)
	repo := NewProdRepository(db)
	mux := http.NewServeMux()
	NewProdHandler(mux, ProdHandlerDeps{
		ProdRepo: repo,
	})
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	srv := http.Server{
		Addr:    ":8088",
		Handler: MidLog(mux),
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
