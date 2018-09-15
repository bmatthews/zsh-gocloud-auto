package generator

import (
	"html/template"
	"io"

	"github.com/bmatthews/zsh-gocloud-auto/completions"
)

var LayoutDir string = "generator"

func Run(s io.Writer, a completions.AutoComplete) {
	tmpl, err := template.ParseGlob(LayoutDir + "/*.gotmpl")
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(s, "compdef", a)
	if err != nil {
		panic(err)
	}
}
