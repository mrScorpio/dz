package validationApi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
	concurrency "github.com/mrScorpio/dz/1-concurrency"
)

type ValHandler struct {
	*VldConfig
}

type ValHandlerDeps struct {
	*VldConfig
}

func NewValHandler(mux *http.ServeMux, deps ValHandlerDeps) {
	handler := &ValHandler{
		VldConfig: deps.VldConfig,
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

		buf, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
		r.Body.Close()
		handler.VldConfig.Address = string(buf)

		hd := HashData{
			Email: handler.VldConfig.Address,
			Hash:  hash,
		}
		if hd.SaveJson() != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		e := email.NewEmail()
		e.To = []string{handler.VldConfig.Address}
		e.Subject = "Verification"
		e.HTML = []byte(fmt.Sprintf("<h1>http://localhost:8086/verify/%s</h1>", hash))
		err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", handler.VldConfig.Email, handler.VldConfig.Password, "smtp.gmail.com"))
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
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
