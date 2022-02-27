package handler

import (
	"encoding/json"
	"net/http"

	"anvarisy.tech/cleangolang/src/modules/auth/endpoint"
)

type Handler struct {
	endpoint endpoint.Endpoint
}

func Init(endpoint endpoint.Endpoint) *Handler {
	return &Handler{
		endpoint: endpoint,
	}
}

func (handler *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var (
		result  endpoint.LoginResponse
		payload endpoint.LoginRequest
		err     error
		ctx     = r.Context()
	)
	var data, _ = json.Marshal(result)
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		endpoint.HTTPStatus(w, http.StatusBadRequest, data)
		return
	}

	result, err = handler.endpoint.Login(ctx, payload)
	if err != nil {
		endpoint.HTTPStatus(w, http.StatusInternalServerError, data)
		return
	}
	endpoint.HTTPStatus(w, 200, data)

}
