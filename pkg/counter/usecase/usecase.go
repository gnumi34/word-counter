package usecase

import (
	"strings"

	"github.com/gnumi34/word-counter/pkg/domain/repository"
	"github.com/gnumi34/word-counter/pkg/domain/response"
	"github.com/gnumi34/word-counter/pkg/utils"
)

type UseCase struct {
	restRepository repository.CounterREST
}

func NewUseCase(restClient repository.CounterREST) *UseCase {
	return &UseCase{
		restRepository: restClient,
	}
}

func (c *UseCase) CountWord(text string) (response.WordCountResult, error) {
	validResult := make(map[string]int)
	invalidResult := make(map[string]int)
	var wordResponse response.EnglishDictionaryResponse
	var err error

	splittedWords := strings.Split(text, " ")
	for idx := range splittedWords {
		processedWord := utils.NonLetterRemover(splittedWords[idx])
		wordResponse, err = c.restRepository.GetWordFromAPI(processedWord)
		if err != nil {
			return response.WordCountResult{}, err
		}

		if len(wordResponse) == 0 {
			if _, ok := invalidResult[splittedWords[idx]]; ok {
				invalidResult[processedWord]++
			} else {
				invalidResult[processedWord] = 1
			}
			continue
		}

		if _, ok := validResult[processedWord]; ok {
			validResult[processedWord]++
		} else {
			validResult[processedWord] = 1
		}
	}

	result := response.WordCountResult{
		WordCount:        validResult,
		InvalidWordCount: invalidResult,
	}

	return result, nil
}
