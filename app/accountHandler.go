package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nk75razor/REST_microservices/dto"
	"github.com/nk75razor/REST_microservices/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())

	} else {
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			WriteResponse(w, appError.Code, appError.Message)
		} else {
			WriteResponse(w, http.StatusCreated, account)
		}
	}
}
