package orderapi

import (
	"encoding/json"
	"log"
	"net/http"
)

type AuthHandlerDeps struct {
	*AuthService
}

type AuthHandler struct {
	*AuthService
}

func NewAuthHandler(mux *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		AuthService: deps.AuthService,
	}
	mux.HandleFunc("POST /auth/login", handler.Login())
	mux.HandleFunc("POST /auth/code", handler.Check())
}

func (handler *AuthHandler) SendSMS(phone string, w http.ResponseWriter) {
	user, err := handler.repo.GetUserByPhone(phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("SMS with code:", user.Code, "is sended to", user.Phone)
}

func JsonResp(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginReq PhoneReq
		err := json.NewDecoder(r.Body).Decode(&loginReq)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ssId, err := handler.AuthService.ChkPhone(loginReq.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		loginResp := SessionResp{
			SessionId: ssId,
		}
		handler.SendSMS(loginReq.Phone, w)
		JsonResp(w, loginResp, http.StatusOK)
	}
}

func (handler *AuthHandler) Check() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var codeReq CodeReq
		err := json.NewDecoder(r.Body).Decode(&codeReq)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := handler.AuthService.ChkCode(codeReq.SessionId, codeReq.Code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if token == "" {
			http.Error(w, "wrong code", http.StatusUnauthorized)
			return
		}
		tokenResp := TokenResp{
			Token: token,
		}
		JsonResp(w, tokenResp, http.StatusOK)
	}
}
