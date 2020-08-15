package restapi

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Code    int32  `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func ResponseBadRequest(msg string, w http.ResponseWriter) {
	resp := APIResponse{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(resp)
}
