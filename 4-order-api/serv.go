package orderapi

import "net/http"

func ProdServ() {
	conf := LoadConfig()
	db := NewDb(conf)
	repo := NewProdRepository(db)
	mux := http.NewServeMux()
	NewProdHandler(mux, ProdHandlerDeps{
		ProdRepo: repo,
	})
	srv := http.Server{
		Addr:    ":8088",
		Handler: mux,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
