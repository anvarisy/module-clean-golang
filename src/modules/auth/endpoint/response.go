package endpoint

import "net/http"

// HTTPStatus responses a http status response for JSONResponse.
func HTTPStatus(w http.ResponseWriter, code int, msg []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(msg)
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredIn    int64  `json:"expired_in"`
}
