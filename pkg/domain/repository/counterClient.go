package repository

import (
	"github.com/gnumi34/word-counter/pkg/domain/response"
)

type CounterREST interface {
	GetWordFromAPI(word string) (response.EnglishDictionaryResponse, error)
}
