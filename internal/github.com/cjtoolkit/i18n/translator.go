package i18n

import (
	"bytes"
	"fmt"
	"text/template"
)

// Translator returns the translation of the string identified by key.
//
// If key is a non-plural form, then the first variadic argument may be a map[string]interface{}
// that contains template data.
//
// If key is a plural form, then the first variadic argument must be an integer type
// (int, int8, int16, int32, int64) or a float formatted as a string (e.g. "123.45").
// The second variadic argument may be a map[string]interface{} that contains template data.
type Translator func(key string, args ...interface{}) string

// Tfunc returns a Translate type that will be bound to the namespace.
func Tfunc(namespace string, languageSources ...string) (Translator, error) {
	for _, languageSource := range languageSources {
		langMutex.RLock()
		l := langMap[languageSource]
		langMutex.RUnlock()
		if l == nil {
			continue
		}
		l.RLock()
		n := l.n[namespace]
		l.RUnlock()
		if n == nil {
			continue
		}
		fn := func(key string, args ...interface{}) string {
			n.RLock()
			value := n.kv[key]
			n.RUnlock()
			if value == nil {
				return ""
			}

			switch value := value.(type) {
			case *template.Template:
				m := map[string]interface{}{}
				if len(args) > 0 {
					if value, ok := args[0].(map[string]interface{}); ok {
						m = value
					}
				}
				buf := &bytes.Buffer{}
				defer buf.Reset()
				err := value.Execute(buf, m)
				if err != nil {
					panic(err)
					return ""
				}
				return buf.String()
			case plural:
				if len(args) < 1 {
					return ""
				}
				m := map[string]interface{}{}
				if len(args) > 1 {
					if value, ok := args[1].(map[string]interface{}); ok {
						m = value
					}
				}
				m["Count"] = args[0]
				buf := &bytes.Buffer{}
				defer buf.Reset()

				ops, err := newOperands(args[0])
				if err != nil {
					panic(err)
					return ""
				}

				switch n.l.pluralFn(*ops) {
				case Invalid:
					return ""
				case Zero:
					if value.Zero == nil {
						return ""
					}
					err = value.Zero.Execute(buf, m)
					if err != nil {
						panic(err)
						return ""
					}
					return buf.String()
				case One:
					if value.One == nil {
						return ""
					}
					err = value.One.Execute(buf, m)
					if err != nil {
						panic(err)
						return ""
					}
					return buf.String()
				case Two:
					if value.Two == nil {
						return ""
					}
					err = value.Two.Execute(buf, m)
					if err != nil {
						panic(err)
						return ""
					}
					return buf.String()
				case Few:
					if value.Few == nil {
						return ""
					}
					err = value.Few.Execute(buf, m)
					if err != nil {
						panic(err)
						return ""
					}
					return buf.String()
				case Many:
					if value.Many == nil {
						return ""
					}
					err = value.Many.Execute(buf, m)
					if err != nil {
						panic(err)
						return ""
					}
					return buf.String()
				case Other:
					if value.Other == nil {
						return ""
					}
					err = value.Other.Execute(buf, m)
					if err != nil {
						panic(err)
						return ""
					}
					return buf.String()
				}
			}

			return ""
		}
		return fn, nil
	}
	return nil, fmt.Errorf("i18n: Unable to verify languageSources or namespace.")
}

// MustTfunc is similar to Tfunc except it panics if an error happens.
func MustTfunc(namespace string, languageSources ...string) Translator {
	fn, err := Tfunc(namespace, languageSources...)
	if err != nil {
		panic(err)
	}
	return fn
}
