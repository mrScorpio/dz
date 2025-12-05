package orderapi

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	StatusWork string = "work"
	StatusTest string = "test"
)

func ProdServ(status string) http.Handler {
	var conf *Config
	if status == StatusWork {
		conf = LoadConfig("")
	}
	if status == StatusTest {
		conf = LoadConfig("../fortest.env")
	}
	db := NewDb(conf)
	repoProd := NewProdRepository(db)
	repoUsers := NewUserRepository(db)
	repoOrders := NewOrderRepository(db)
	mux := http.NewServeMux()
	NewProdHandler(mux, ProdHandlerDeps{
		ProdRepo: repoProd,
		UserRepo: repoUsers,
	})
	NewOrderHandler(mux, OrderHandlerDeps{
		ProdRepo:  repoProd,
		UserRepo:  repoUsers,
		OrderRepo: repoOrders,
	})
	authService := NewAuthService(repoUsers)
	NewAuthHandler(mux, AuthHandlerDeps{
		AuthService: authService,
	})
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	return MidLog(mux)
}
