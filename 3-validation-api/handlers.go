package validationApi

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"

	"github.com/jordan-wright/email"
	concurrency "github.com/mrScorpio/dz/1-concurrency"
)

type ValHandler struct {
	*Config
}

type ValHandlerDeps struct {
	*Config
}

func NewValHandler(mux *http.ServeMux, deps ValHandlerDeps) {
	handler := &ValHandler{
		Config: deps.Config,
	}
	mux.HandleFunc("POST /send", handler.Send())
	mux.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *ValHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hashSl := concurrency.GenNums(6, 26)
		hashR := make([]rune, len(hashSl))
		for i, v := range hashSl {
			hashR[i] = rune(v + 97)
		}
		hash := string(hashR)
		buf := []byte{}
		_, err := r.Body.Read(buf)
		defer r.Body.Close()

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
		handler.Config.Address = string(buf)

		e := email.NewEmail()
		e.To = []string{handler.Config.Address}
		e.Subject = "Verification"
		e.HTML = []byte(fmt.Sprintf("<h1>http://localhost:8086/verify/%s</h1>", hash))
		err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", handler.Config.Email, handler.Config.Password, "smtp.gmail.com"))
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		hd := HashData{
			Email: handler.Config.Address,
			Hash:  hash,
		}
		if hd.SaveJson() != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (handler *ValHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		hash := strings.TrimPrefix(r.URL.Path, "/verify/")
		var hd HashData
		if hd.ReadJson() != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		if os.Remove("hashfile.json") != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		if hash == hd.Hash {
			w.WriteHeader(http.StatusAccepted)
			_, err := w.Write([]byte("Your email is verified!"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			w.WriteHeader(http.StatusForbidden)
		}

	}
}
