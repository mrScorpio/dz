package orderapi

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func ProdServ() {
	conf := LoadConfig()
	db := NewDb(conf)
	repoProd := NewProdRepository(db)
	repoUsers := NewUserRepository(db)
	mux := http.NewServeMux()
	NewProdHandler(mux, ProdHandlerDeps{
		ProdRepo: repoProd,
	})
	authService := NewAuthService(repoUsers)
	NewAuthHandler(mux, AuthHandlerDeps{
		AuthService: authService,
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
