package resources

import (
	"errors"
	"strings"

	"github.com/dragonator/gopher-translator/internal/service/svc"
	"github.com/dragonator/gopher-translator/internal/storage"
	"github.com/dragonator/gopher-translator/internal/translator"
)

// Gopher -
type Gopher interface {
	TranslateWord(word string) (string, error)
	TranslateSentence(sentence string) (string, error)
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
func (gr *gopher) TranslateWord(word string) (string, error) {
	translation, err := gr.translator.Translate(word)
	if err != nil {
		return "", err
	}
	gr.store.AddRecord(&storage.Record{Input: word, Output: translation})
	return translation, nil
}

// TranslateSentence -
func (gr *gopher) TranslateSentence(sentence string) (string, error) {
	endSymbol := string(sentence[len(sentence)-1])
	words := strings.Split(sentence[:len(sentence)-1], " ")

	translatedWords := make([]string, 0, len(words)+1)
	for _, w := range words {
		tw, err := gr.translator.Translate(w)
		if err != nil {
			if errors.Is(err, svc.ErrInvalidInput) {
				// skip invalid words in translated sentence
				continue
			}
			return "", err
		}
		translatedWords = append(translatedWords, tw)
	}
	translatedSentence := strings.Join(translatedWords, " ") + endSymbol

	gr.store.AddRecord(&storage.Record{Input: sentence, Output: translatedSentence})
	return translatedSentence, nil
}

// History -
func (gr *gopher) History() (result []map[string]string) {
	for _, r := range gr.store.History() {
		result = append(result, map[string]string{r.Input: r.Output})
	}
	return
}
