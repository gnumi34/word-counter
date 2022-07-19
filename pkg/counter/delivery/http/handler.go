package http

import (
	"encoding/json"
	"log"
	"net/http"

	commonHTTP "github.com/gnumi34/word-counter/pkg/common/delivery/http"
	"github.com/gnumi34/word-counter/pkg/counter/usecase"
	"github.com/gnumi34/word-counter/pkg/domain/request"
	"github.com/gnumi34/word-counter/pkg/domain/response"
)

type Handler struct {
	commonHTTP.CommonHandler
	useCase *usecase.UseCase
}

func NewHandler(usecase *usecase.UseCase) *Handler {
	return &Handler{
		useCase: usecase,
	}
}

func (h *Handler) CountWordsFromText(w http.ResponseWriter, r *http.Request) {
	var textReq request.TextRequest
	var result response.WordCountResult
	var err error

	w.Header().Add("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		resp := response.CommonResponse{
			Message: response.DataFailed,
			Error:   "Method is not allowed",
		}
		respBytes, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(respBytes)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&textReq)
	if err != nil {
		log.Println(err.Error())
		resp := response.CommonResponse{
			Message: response.DataFailed,
			Error:   response.BadRequest,
		}
		respBytes, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(respBytes)
		return
	}

	result, err = h.useCase.CountWord(textReq.Input)
	if err != nil {
		log.Println(err.Error())
		resp := response.CommonResponse{
			Message: response.DataFailed,
			Error:   response.UnknownError,
		}
		respBytes, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusNotFound)
		w.Write(respBytes)
		return
	}

	resp := response.CommonResponse{
		Message: response.DataSuccess,
		Data:    result,
	}
	respBytes, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
	return
}
