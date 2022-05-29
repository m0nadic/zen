package util

import (
	"io"
	"text/template"
	"zen/internal/pkg/functions"
)

func ExpandTemplate(fileName string, w io.Writer, data any) error {
	tmpl, err := template.New(fileName).
		Funcs(functions.GetFuncMap()).
		ParseFiles(fileName)

	if err != nil {
		return err
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
