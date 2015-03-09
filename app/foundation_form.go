package app

import (
	"github.com/cjtoolkit/form"
)

type foundationForm struct {
	Title      string
	Firstname  string
	Secondname string
}

func (b *foundationForm) CJForm(f *form.Fields) {
	// Title
	func() {
		f := f.Init("Title", form.Select)
		f.Options([]form.Option{
			{Value: "Mr.", Content: "Mr.", Selected: true},
			{Value: "Mrs.", Content: "Mrs."},
		})
		f.Mandatory()

		html := f.HTML()
		html.Before = `<div class="row"><div class="large-12 columns">`
		html.After = `</div></div>`
	}()

	// Firstname
	func() {
		f := f.Init("Firstname", form.InputText)
		f.Mandatory()

		size := f.Size()
		size.Max = 32

		f.Attr(map[string]string{
			"placeholder": "First Name",
		})

		html := f.HTML()
		html.Before = `<div class="row"><div class="large-6 columns">`
		html.After = `</div>`
	}()

	// Secondname
	func() {
		f := f.Init("Secondname", form.InputText)
		f.Mandatory()

		size := f.Size()
		size.Max = 32

		f.Attr(map[string]string{
			"placeholder": "First Name",
		})

		html := f.HTML()
		html.Before = `<div class="large-6 columns">`
		html.After = `</div></div>`
	}()
}
