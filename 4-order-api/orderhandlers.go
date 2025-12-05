package orderapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	ProdRepo  *ProdRepository
	UserRepo  *UserRepository
	OrderRepo *OrderRepository
}

type OrderHandlerDeps struct {
	ProdRepo  *ProdRepository
	UserRepo  *UserRepository
	OrderRepo *OrderRepository
}

func NewOrderHandler(mux *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{
		ProdRepo:  deps.ProdRepo,
		UserRepo:  deps.UserRepo,
		OrderRepo: deps.OrderRepo,
	}
	mux.Handle("POST /order", IsAuthed(handler.Create(), handler.UserRepo))
	mux.Handle("GET /order/{id}", IsAuthed(handler.GetById(), handler.UserRepo))
	mux.Handle("GET /my-orders", IsAuthed(handler.GetAll(), handler.UserRepo))
}

func (h *OrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newReq NewOrderReq
		w.Header().Set("Content-Type", "text/plain")
		if err := json.NewDecoder(r.Body).Decode(&newReq); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		trueProdSl := []uint{}
		for _, v := range newReq.ProductIds {
			_, err := h.ProdRepo.GetById(v)
			if err == nil {
				trueProdSl = append(trueProdSl, v)
			} else {
				w.Write([]byte("product id " + fmt.Sprint(v) + " is ignored because is not exist in menu\n"))
			}
		}

		phone := r.Context().Value(CtxPhoneKey).(string)
		user, err := h.UserRepo.GetUserByPhone(phone)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		newOrd := NewOrder(user.ID, newReq.Address, trueProdSl)
		storedOrder, err := h.OrderRepo.Create(newOrd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error: " + err.Error()))
			return
		}

		w.Write([]byte(fmt.Sprintf("New order id=%d with %d items is created", storedOrder.ID, len(storedOrder.Products))))
	}
}

func (h *OrderHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		phone := r.Context().Value(CtxPhoneKey).(string)
		user, err := h.UserRepo.GetUserByPhone(phone)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		order, err := h.OrderRepo.GetById(uint(id), user.ID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("error: " + err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(order); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("error: " + err.Error()))
			return
		}
	}
}

func (h *OrderHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		phone := r.Context().Value(CtxPhoneKey).(string)
		user, err := h.UserRepo.GetUserByPhone(phone)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error: " + err.Error()))
			return
		}

		orders, err := h.OrderRepo.GetByUser(user.ID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("error: " + err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(orders); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("error: " + err.Error()))
			return
		}
	}
}
