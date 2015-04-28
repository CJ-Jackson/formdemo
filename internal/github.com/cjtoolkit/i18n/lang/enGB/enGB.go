// English (British)
package enGB

import (
	"github.com/CJ-Jackson/formdemo/internal/github.com/cjtoolkit/i18n"
	"github.com/CJ-Jackson/formdemo/internal/github.com/cjtoolkit/i18n/plural/en"
)

func init() {
	i18n.InitLanguage("en-GB", en.Plural)
	i18n.Alias("en_GB", "en-GB")
	i18n.Alias("enGB", "en-GB")
}
