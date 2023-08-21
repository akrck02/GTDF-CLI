package core

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/akrck02/GTDF-CLI/api/io"
	"github.com/akrck02/GTDF-CLI/api/logger"
	"github.com/akrck02/GTDF-CLI/api/templates"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func NewView(name string) {

	logger.Log("Current directory: " + io.GetCurrentDirectory())
	if !IsProject(io.GetCurrentDirectory()) {
		logger.Error("You are not in a GTDF project directory.")
		return
	}

	logger.Log("Generating view " + name + "...")
	caser := cases.Title(language.English)
	name = caser.String(name)

	io.SecureMkdir(io.GetCurrentDirectory()+"/frontend/src/views/"+strings.ToLower(name), 0755)

	template := viewTemplate(name)
	coreTemplate := viewCoreTemplate(name)
	//logger.Log(template)

	io.WriteFileWith(io.GetCurrentDirectory()+"/frontend/src/views/"+strings.ToLower(name)+"/"+name+".ts", template)
	io.WriteFileWith(io.GetCurrentDirectory()+"/frontend/src/views/"+strings.ToLower(name)+"/"+name+".core.ts", coreTemplate)

	logger.Log("View generated.")
}

func viewTemplate(name string) string {
	templ := template.Must(template.New("view").Parse(templates.VIEW))
	var tpl bytes.Buffer

	templ.Execute(&tpl, map[string]interface{}{
		"name": name,
	})

	return tpl.String()
}

func viewCoreTemplate(name string) string {

	templ := template.Must(template.New("view").Parse(templates.VIEW_CORE))
	var tpl bytes.Buffer

	templ.Execute(&tpl, map[string]interface{}{
		"name": name,
	})

	return tpl.String()

}
