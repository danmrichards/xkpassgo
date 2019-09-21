package generator

type Config struct {
	NumWords   int `json:"num_words"`
	WordLenMin int `json:"word_length_min"`
	WordLenMax int `json:"word_length_max"`
}
