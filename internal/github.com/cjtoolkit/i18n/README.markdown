# CJToolkit i18n

i18n is a Go package originally based on (and inspired by) [Nick Snyder's go-i18n](https://github.com/nicksnyder/go-i18n), it is more geared up towards hard-coded translation rather than using translation files.
* Supports pluralized strings using [CLDR plural rules](http://cldr.unicode.org/index/cldr-spec/plural-rules).
* Supports strings with named variables using [text/template](http://golang.org/pkg/text/template/) syntax.
* Each Language and Dialect is in it own designated package (located in 'lang'), only need to import the Languages and Dialects you going to need in your project or product. (There is no default language!)
* Supports namespace.
* Supports language alias.

Documentation can be found at.

https://godoc.org/github.com/cjtoolkit/i18n

## Installation

~~~
go get github.com/cjtoolkit/i18n
~~~

## Usage

~~~ go
package main

import (
	"fmt"
	"github.com/cjtoolkit/i18n"

	// Init English (British)
	_ "github.com/cjtoolkit/i18n/lang/enGB"
)

func init() {
	// Language, Namespace, Map
	i18n.AppendMap("en-GB", "example", map[string]interface{}{
		"hello": "Hello, {{.Person}}!",
		"apple": i18n.Plural{
			One:   "There is an Apple!",
			Other: "There is {{.Count}} Apples!",
		},
	})
}

func main() {
	// Namespace, Language Sources
	T := i18n.MustTfunc("example", "en-GB")

	fmt.Println(T("hello", map[string]interface{}{
		"Person": "Chris",
	}))

	fmt.Println(T("apple", 1))

	fmt.Println(T("apple", 50))
}
~~~

## Adding New Languages

It is easy to add support for additional languages and dialects:

1. Lookup the language's [CLDR plural rules](http://www.unicode.org/cldr/charts/latest/supplemental/language_plural_rules.html).

2. Create a new package in 'lang' (and 'plural' if necessary).  Use existing package in 'lang' and 'plural' as examples.

3. Make sure it has a test case.

4. Submit a pull request to towards 'develop' branch.

## Buy me a beer!

Bitcoin - 1MieXR5ANYY6VstNanhuLRtGQGn6zpjxK3