package functions

import (
	"text/template"
)

func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"UUID": UUID,
		"seq":  seq,
		"fake": fakeNS,
	}
}
