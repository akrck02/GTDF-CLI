package core

import (
	"bytes"
	"html/template"

	"github.com/akrck02/GTDF-CLI/api/io"
	"github.com/akrck02/GTDF-CLI/api/logger"
	"github.com/akrck02/GTDF-CLI/api/templates"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func NewService(name string) {

	logger.Log("Current directory: " + io.GetCurrentDirectory())
	if !IsProject(io.GetCurrentDirectory()) {
		logger.Error("You are not in a GTDF project directory.")
		return
	}

	logger.Log("Generating service " + name + "...")
	caser := cases.Title(language.English)
	name = caser.String(name)

	template := serviceTemplate(name)
	io.WriteFileWith(io.GetCurrentDirectory()+"/frontend/src/services/"+name+".ts", template)

	logger.Log("Service generated.")
}

func serviceTemplate(name string) string {
	templ := template.Must(template.New("view").Parse(templates.SERVICE))
	var tpl bytes.Buffer

	templ.Execute(&tpl, map[string]interface{}{
		"name": name,
	})

	return tpl.String()
}
