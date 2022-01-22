package resources

import (
	"strings"

	"github.com/dragonator/gopher-translator/internal/storage"
	"github.com/dragonator/gopher-translator/internal/translator"
)

// Gopher -
type Gopher interface {
	TranslateWord(word string) string
	TranslateSentence(sentence string) string
	History() []map[string]string
}

type gopher struct {
	translator translator.Translator
	store      storage.Storage
}

// NewGopher -
func NewGopher(tr translator.Translator, st storage.Storage) Gopher {
	return &gopher{
		translator: tr,
		store:      st,
	}
}

// TranslateWord -
func (gr *gopher) TranslateWord(word string) string {
	translation := gr.translator.Translate(word)
	gr.store.AddRecord(&storage.Record{Input: word, Output: translation})
	return translation
}

// TranslateSentence -
func (gr *gopher) TranslateSentence(sentence string) string {
	endSymbol := string(sentence[len(sentence)-1])
	words := strings.Split(sentence[:len(sentence)-2], " ")

	translatedWords := make([]string, 0, len(words)+1)
	for _, w := range words {
		tw := gr.translator.Translate(w)
		translatedWords = append(translatedWords, tw)
	}
	translatedSentence := strings.Join(translatedWords, " ") + endSymbol

	gr.store.AddRecord(&storage.Record{Input: sentence, Output: translatedSentence})
	return translatedSentence
}

// History -
func (gr *gopher) History() (result []map[string]string) {
	for _, r := range gr.store.History() {
		result = append(result, map[string]string{r.Input: r.Output})
	}
	return
}
