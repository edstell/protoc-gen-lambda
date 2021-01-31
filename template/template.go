package template

import (
	"fmt"
	"text/template"
)

// FromAsset initializes a template from the named asset.
func FromAsset(name string, funcs map[string]interface{}) (*template.Template, error) {
	data, err := Asset(fmt.Sprintf("%s.gotmpl", name))
	if err != nil {
		return nil, err
	}
	t, err := template.New(name).Funcs(funcs).Parse(string(data))
	if err != nil {
		return nil, err
	}
	return t, nil
}
