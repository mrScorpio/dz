package validationApi

import (
	"log"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
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
	mux.HandleFunc("GET /verify/", handler.Verify())
}

func (handler *ValHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e := email.NewEmail()
		e.To = []string{handler.Address}
		e.Subject = "Verification"
		e.HTML = []byte("<h1>http://localhost:8086/verify/{hash}</h1>")
		err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", handler.Email, handler.Password, "smtp.gmail.com"))
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (handler *ValHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		hash := strings.TrimPrefix(r.URL.Path, "/verify/")
		if hash == "{hash}" {
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
