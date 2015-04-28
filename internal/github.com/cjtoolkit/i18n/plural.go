package i18n

import (
	"text/template"
)

// PluralCode represents a language pluralization form as defined here:
// http://cldr.unicode.org/index/cldr-spec/plural-rules
type PluralCode uint8

// All defined plural categories
const (
	Invalid PluralCode = iota
	Zero
	One
	Two
	Few
	Many
	Other
)

// Specify Plural.
type Plural struct {
	Zero, One, Two, Few, Many, Other string
}

type plural struct {
	Zero, One, Two, Few, Many, Other *template.Template
}
