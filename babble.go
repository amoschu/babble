package babble

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Babbler struct {
	Separator  string   // Babble() word separator
	Words      []string // List of possible words (set by NewBabbler())
}

// NewBabbler initializes a new Babbler instance.
func NewBabbler() (b Babbler) {
	b.Separator = "-"
	b.Words = readAvailableDictionary()
	return
}

// Babble returns a random dictionary word.
func (this Babbler) Babble(c int) string {
	if c <= 0 {
		return ""
	}
	pieces := make([]string, c)
	for i := 0; i < c; i++ {
		word := this.Words[rand.Int()%len(this.Words)]
		word = strings.Map(func(r rune) rune {
			if unicode.IsPunct(r) {
				return -1
			}
			return r
		}, word)
		pieces[i] = word
	}

	return strings.Join(pieces, this.Separator)
}
