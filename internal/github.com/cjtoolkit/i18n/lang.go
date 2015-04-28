package i18n

import (
	"sync"
)

type lang struct {
	sync.RWMutex
	pluralFn func(Operands) PluralCode
	n        map[string]*ns
}

var langMutex sync.RWMutex
var langMap = map[string]*lang{}

// Initialise Language.
func InitLanguage(langCode string, pluralFn func(Operands) PluralCode) {
	langMutex.Lock()
	defer langMutex.Unlock()
	langMap[langCode] = &lang{
		pluralFn: pluralFn,
		n:        map[string]*ns{},
	}
}

// Specify Alias of a Language.
func Alias(langCode, of string) {
	langMutex.Lock()
	defer langMutex.Unlock()
	langMap[langCode] = langMap[of]
}
