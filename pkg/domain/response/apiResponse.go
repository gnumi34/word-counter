package response

const (
	DataSuccess  = "Successfully fetched data"
	DataFailed   = "Failed to fetch data"
	BadRequest   = "Bad Request. Try again."
	UnknownError = "Something is wrong. Please try again later."
)

type CommonResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

type WordCountResult struct {
	WordCount        map[string]int `json:"word_count"`
	InvalidWordCount map[string]int `json:"invalid_word_count"`
}

type License struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Phonetic struct {
	Audio     string  `json:"audio"`
	SourceURL string  `json:"sourceUrl,omitempty"`
	License   License `json:"license,omitempty"`
	Text      string  `json:"text,omitempty"`
}

type Definition struct {
	Definition string        `json:"definition"`
	Synonyms   []interface{} `json:"synonyms"`
	Antonyms   []interface{} `json:"antonyms"`
}

type Meaning struct {
	PartOfSpeech string        `json:"partOfSpeech"`
	Definitions  []Definition  `json:"definitions"`
	Synonyms     []string      `json:"synonyms"`
	Antonyms     []interface{} `json:"antonyms"`
}

type EnglishDictionaryResponse []struct {
	Word       string     `json:"word"`
	Phonetics  []Phonetic `json:"phonetics"`
	Meanings   []Meaning  `json:"meanings"`
	License    License    `json:"license"`
	SourceUrls []string   `json:"sourceUrls"`
}
