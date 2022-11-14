package generator

import (
	"bytes"
	"fmt"
	"github.com/sasswart/gin-in-a-can/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"text/template"
)

// Generate is the main parsing and rendering steps within the generator library
func Generate(config config.Config, templateFile string) ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})

	templater := template.New(templateFile)

	templater.Funcs(templateFuncMap)

	parsedTemplate, err := templater.ParseGlob(fmt.Sprintf("%s/*.tmpl", config.TemplateDirectory))
	if err != nil {
		return nil, err
	}

	err = parsedTemplate.Execute(buff, config)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

var templateFuncMap = template.FuncMap{
	"ToUpper":  strings.ToUpper,
	"ToTitle":  toTitleCase,
	"Type":     Type,
	"Sanitize": Sanitize,
}

func toTitleCase(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}
