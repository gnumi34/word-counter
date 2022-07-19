package http

import (
	"encoding/json"
	"net/http"

	"github.com/gnumi34/word-counter/pkg/domain/response"
)

type CommonHandler struct{}

func (c *CommonHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	resp := response.CommonResponse{
		Message: "Hello World! System is functioning just fine!",
	}

	respBytes, _ := json.Marshal(resp)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}
