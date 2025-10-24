package randomApi

import (
	"fmt"
	"net/http"

	concurrency "github.com/mrScorpio/dz/1-concurrency"
)

func sendRandNum(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprint(concurrency.GenNums(1, 6)[0])))
}

func crWebServ() {
	mux := http.NewServeMux()
	mux.HandleFunc("/randnum", sendRandNum)
	serv := http.Server{
		Addr:    ":8086",
		Handler: mux,
	}
	panic(serv.ListenAndServe())
}
