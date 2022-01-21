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
	replacer   *strings.Replacer
}

// NewGopher -
func NewGopher(tr translator.Translator, st storage.Storage) Gopher {
	return &gopher{
		translator: tr,
		store:      st,
		replacer:   strings.NewReplacer("'", "", "’", ""),
	}
}

// TranslateWord -
func (gr *gopher) TranslateWord(word string) string {
	normalizedWord := gr.replacer.Replace(word)
	translation := gr.translator.Translate(normalizedWord)
	gr.store.AddRecord(&storage.Record{Input: word, Output: translation})
	return translation
}

// TranslateSentence -
func (gr *gopher) TranslateSentence(sentence string) string {
	endSymbol := string(sentence[len(sentence)-1])
	words := strings.Split(sentence[:len(sentence)-2], " ")

	translatedWords := make([]string, 0, len(words)+1)
	for _, w := range words {
		normalizedWord := gr.replacer.Replace(w)
		tw := gr.translator.Translate(normalizedWord)
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

func normalizeWord(word string) {

}
