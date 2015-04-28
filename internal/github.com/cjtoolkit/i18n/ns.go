package i18n

import (
	"fmt"
	"sync"
	"text/template"
)

type ns struct {
	sync.RWMutex
	l  *lang
	kv map[string]interface{}
}

// Set Key and Value to namespace and langCode.
func SetKeyValue(langCode, namespace, key string, value interface{}) {
	if value == nil {
		return
	}
	langMutex.RLock()
	if langMap[langCode] == nil {
		langMutex.RUnlock()
		panic(fmt.Errorf("i18n: langCode '%s' has not been initialise"))
		return
	}
	l := langMap[langCode]
	langMutex.RUnlock()
	l.Lock()
	if l.n[namespace] == nil {
		l.n[namespace] = &ns{
			l:  l,
			kv: map[string]interface{}{},
		}
	}
	n := l.n[namespace]
	l.Unlock()
	n.Lock()
	defer n.Unlock()
	switch value := value.(type) {
	case string:
		t, err := template.New("").Parse(value)
		if err != nil {
			panic(err)
		}
		n.kv[key] = t
	case Plural:
		pl := &plural{}
		if value.Zero != "" {
			t, err := template.New("").Parse(value.Zero)
			if err != nil {
				panic(err)
			}
			pl.Zero = t
		}
		if value.One != "" {
			t, err := template.New("").Parse(value.One)
			if err != nil {
				panic(err)
			}
			pl.One = t
		}
		if value.Two != "" {
			t, err := template.New("").Parse(value.Two)
			if err != nil {
				panic(err)
			}
			pl.Two = t
		}
		if value.Few != "" {
			t, err := template.New("").Parse(value.Few)
			if err != nil {
				panic(err)
			}
			pl.Few = t
		}
		if value.Many != "" {
			t, err := template.New("").Parse(value.Many)
			if err != nil {
				panic(err)
			}
			pl.Many = t
		}
		if value.Other != "" {
			t, err := template.New("").Parse(value.Other)
			if err != nil {
				panic(err)
			}
			pl.Other = t
		}
		n.kv[key] = *pl
	}
}

// Append Map to namespace and language code.
func AppendMap(langCode, namespace string, m map[string]interface{}) {
	for key, value := range m {
		SetKeyValue(langCode, namespace, key, value)
	}
}
