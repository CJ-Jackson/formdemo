package enGB

import (
	"github.com/CJ-Jackson/formdemo/internal/github.com/cjtoolkit/i18n"
	"testing"
)

func TestEnGB(t *testing.T) {
	i18n.SetKeyValue("en-GB", "test", "person", "Hello, {{.Person}}!")

	T := i18n.MustTfunc("test", "en-GB")

	if value := T("person", map[string]interface{}{
		"Person": "Chris",
	}); value != "Hello, Chris!" {
		t.Errorf("Expected output was 'Hello, Chris!' actual output was '%s'", value)
	}

	i18n.SetKeyValue("en-GB", "test", "apple", i18n.Plural{
		One:   "There is an Apple!",
		Other: "There is {{.Count}} Apples!",
	})

	if value := T("apple", 1); value != "There is an Apple!" {
		t.Errorf("Expected output was 'There is an Apple!' actual output was '%s'", value)
	}

	if value := T("apple", 50); value != "There is 50 Apples!" {
		t.Errorf("Expected output was 'There is 50 Apples!' actual output was '%s'", value)
	}
}
