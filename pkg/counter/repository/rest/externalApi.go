package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gnumi34/word-counter/pkg/domain/response"
	"github.com/gnumi34/word-counter/pkg/utils"
)

type RESTClient struct {
	client *http.Client
}

func NewRESTClient(client *http.Client) *RESTClient {
	return &RESTClient{
		client: client,
	}
}

func (r *RESTClient) GetWordFromAPI(word string) (response.EnglishDictionaryResponse, error) {
	var result response.EnglishDictionaryResponse
	var req *http.Request
	var resp *http.Response
	var err error

	link := "https://api.dictionaryapi.dev/api/v2/entries/en/" + word
	req, err = http.NewRequest(http.MethodGet, link, http.NoBody)
	if err != nil {
		return response.EnglishDictionaryResponse{}, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err = r.client.Do(req)
	if err != nil {
		return response.EnglishDictionaryResponse{}, err
	}
	defer r.client.CloseIdleConnections()

	if resp.StatusCode != http.StatusOK {
		return response.EnglishDictionaryResponse{}, nil
	}

	bodyData := utils.ResponseHttpBody(resp.Body)
	err = json.Unmarshal(bodyData, &result)
	if err != nil {
		return response.EnglishDictionaryResponse{}, err
	}
	return result, nil
}
