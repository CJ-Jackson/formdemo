package en

import (
	"github.com/CJ-Jackson/formdemo/internal/github.com/cjtoolkit/i18n"
)

// Plural rules for English Language
func Plural(ops i18n.Operands) i18n.PluralCode {
	if ops.I == 1 && ops.V == 0 {
		return i18n.One
	}
	return i18n.Other
}
