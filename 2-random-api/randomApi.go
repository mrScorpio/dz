package randomApi

import (
	"fmt"
	"log"
	"net/http"

	concurrency "github.com/mrScorpio/dz/1-concurrency"
)

func sendRandNum(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprint(1 + concurrency.GenNums(1, 6)[0])))
}

func CrWebServ() {
	mux := http.NewServeMux()
	mux.HandleFunc("/randnum", sendRandNum)
	serv := http.Server{
		Addr:    ":8086",
		Handler: mux,
	}
	log.Fatal((serv.ListenAndServe()))
}
