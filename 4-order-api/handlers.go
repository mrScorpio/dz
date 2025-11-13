package orderapi

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ProdHandler struct {
	ProdRepo *ProdRepository
}

type ProdHandlerDeps struct {
	ProdRepo *ProdRepository
}

func NewProdHandler(mux *http.ServeMux, deps ProdHandlerDeps) {
	handler := &ProdHandler{
		ProdRepo: deps.ProdRepo,
	}
	mux.HandleFunc("POST /product", handler.Create())
	mux.HandleFunc("PATCH /product/{id}", handler.Update())
	mux.HandleFunc("DELETE /product/{id}", handler.Delete())
	mux.HandleFunc("GET /product/{id}", handler.Show())
	mux.HandleFunc("GET /products", handler.ShowAll())
}

func (h *ProdHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newProd Product
		w.Header().Set("Content-Type", "text/plain")
		if err := json.NewDecoder(r.Body).Decode(&newProd); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		storedProd, err := h.ProdRepo.Create(&newProd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(storedProd); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("error: " + err.Error()))
		}
	}
}

func (h *ProdHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var editProd Product
		w.Header().Set("Content-Type", "text/plain")
		if err := json.NewDecoder(r.Body).Decode(&editProd); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		editProd.ID = uint(id)
		storedProd, err := h.ProdRepo.Update(&editProd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(storedProd); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("error: " + err.Error()))
		}
	}
}

func (h *ProdHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		_, err = h.ProdRepo.GetById(uint(id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		if err := h.ProdRepo.Delete(uint(id)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (h *ProdHandler) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		foundProd, err := h.ProdRepo.GetById(uint(id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(foundProd); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("error: " + err.Error()))
			return
		}
		//w.WriteHeader(http.StatusOK)
	}
}

func (h *ProdHandler) ShowAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		prods, err := h.ProdRepo.GetSlice(11)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(prods); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("error: " + err.Error()))
			return
		}
		//w.WriteHeader(http.StatusOK)
	}
}
