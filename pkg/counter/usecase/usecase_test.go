package usecase_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/gnumi34/word-counter/pkg/counter/usecase"
	"github.com/gnumi34/word-counter/pkg/domain/mocks"
	"github.com/gnumi34/word-counter/pkg/domain/response"
)

type wordMock struct {
	word   string
	output response.EnglishDictionaryResponse
	err    error
}

type testTable struct {
	name       string
	input      string
	want       response.WordCountResult
	mockOutput []wordMock
	wantErr    bool
}

var tests = []testTable{
	{
		name:  "normal case",
		input: "Go is a statically typed",
		want: response.WordCountResult{
			WordCount: map[string]int{
				"go":    1,
				"is":    1,
				"a":     1,
				"typed": 1,
			},
			InvalidWordCount: map[string]int{
				"statically": 1,
			},
		},
		mockOutput: []wordMock{
			{
				word: "go",
				output: response.EnglishDictionaryResponse{
					{
						Word:       "go",
						Phonetics:  []response.Phonetic{},
						Meanings:   []response.Meaning{},
						License:    response.License{},
						SourceUrls: []string{},
					},
				},
				err: nil,
			},
			{
				word: "is",
				output: response.EnglishDictionaryResponse{
					{
						Word:       "is",
						Phonetics:  []response.Phonetic{},
						Meanings:   []response.Meaning{},
						License:    response.License{},
						SourceUrls: []string{},
					},
				},
				err: nil,
			},
			{
				word: "a",
				output: response.EnglishDictionaryResponse{
					{
						Word:       "a",
						Phonetics:  []response.Phonetic{},
						Meanings:   []response.Meaning{},
						License:    response.License{},
						SourceUrls: []string{},
					},
				},
				err: nil,
			},
			{
				word: "typed",
				output: response.EnglishDictionaryResponse{
					{
						Word:       "typed",
						Phonetics:  []response.Phonetic{},
						Meanings:   []response.Meaning{},
						License:    response.License{},
						SourceUrls: []string{},
					},
				},
				err: nil,
			},
			{
				word:   "statically",
				output: response.EnglishDictionaryResponse{},
				err:    nil,
			},
		},
		wantErr: false,
	},
	{
		name:  "non-letter test",
		input: "Go. i13s a stat1ic4ally@ typed!",
		want: response.WordCountResult{
			WordCount: map[string]int{
				"go":    1,
				"is":    1,
				"a":     1,
				"typed": 1,
			},
			InvalidWordCount: map[string]int{
				"statically": 1,
			},
		},
		mockOutput: []wordMock{
			{
				word: "go",
				output: response.EnglishDictionaryResponse{
					{
						Word:       "go",
						Phonetics:  []response.Phonetic{},
						Meanings:   []response.Meaning{},
						License:    response.License{},
						SourceUrls: []string{},
					},
				},
				err: nil,
			},
			{
				word: "is",
				output: response.EnglishDictionaryResponse{
					{
						Word:       "is",
						Phonetics:  []response.Phonetic{},
						Meanings:   []response.Meaning{},
						License:    response.License{},
						SourceUrls: []string{},
					},
				},
				err: nil,
			},
			{
				word: "a",
				output: response.EnglishDictionaryResponse{
					{
						Word:       "a",
						Phonetics:  []response.Phonetic{},
						Meanings:   []response.Meaning{},
						License:    response.License{},
						SourceUrls: []string{},
					},
				},
				err: nil,
			},
			{
				word: "typed",
				output: response.EnglishDictionaryResponse{
					{
						Word:       "typed",
						Phonetics:  []response.Phonetic{},
						Meanings:   []response.Meaning{},
						License:    response.License{},
						SourceUrls: []string{},
					},
				},
				err: nil,
			},
			{
				word:   "statically",
				output: response.EnglishDictionaryResponse{},
				err:    nil,
			},
		},
		wantErr: false,
	},
	{
		name:  "external timeout error",
		input: "Go is a statically typed",
		want:  response.WordCountResult{},
		mockOutput: []wordMock{
			{
				word: "go",
				output: response.EnglishDictionaryResponse{
					{
						Word:       "go",
						Phonetics:  []response.Phonetic{},
						Meanings:   []response.Meaning{},
						License:    response.License{},
						SourceUrls: []string{},
					},
				},
				err: nil,
			},
			{
				word: "is",
				output: response.EnglishDictionaryResponse{
					{
						Word:       "is",
						Phonetics:  []response.Phonetic{},
						Meanings:   []response.Meaning{},
						License:    response.License{},
						SourceUrls: []string{},
					},
				},
				err: nil,
			},
			{
				word: "a",
				output: response.EnglishDictionaryResponse{
					{
						Word:       "a",
						Phonetics:  []response.Phonetic{},
						Meanings:   []response.Meaning{},
						License:    response.License{},
						SourceUrls: []string{},
					},
				},
				err: errors.New("context timeout deadline exceeded"),
			},
		},
		wantErr: true,
	},
}

func TestUseCase_CountWord(t *testing.T) {
	mockCounterRestRepository := mocks.NewCounterREST(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := usecase.NewUseCase(mockCounterRestRepository)
			for idx := range tt.mockOutput {
				mockCounterRestRepository.On("GetWordFromAPI", tt.mockOutput[idx].word).Return(tt.mockOutput[idx].output, tt.mockOutput[idx].err).Once()
			}
			got, err := c.CountWord(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.CountWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.CountWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
