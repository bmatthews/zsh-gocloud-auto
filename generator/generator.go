package generator

import (
	"html/template"
	"io"
	"zsh-go-auto/completions"
)

var LayoutDir string = "generator"

func Run(s io.Writer, a *completions.AutoComplete) {
	tmpl, err := template.ParseGlob(LayoutDir + "/*.gotmpl")
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(s, "compdef", a)
	if err != nil {
		panic(err)
	}
}
