package core

import (
	"bytes"

	"github.com/akrck02/GTDF-CLI/api/io"
	"github.com/akrck02/GTDF-CLI/api/logger"
	"github.com/akrck02/GTDF-CLI/api/templates"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func NewService(name string) {

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

func DeleteService(name string) {

	if !IsProject(io.GetCurrentDirectory()) {
		logger.Error("You are not in a GTDF project directory.")
		return
	}

	logger.Log("Deleting service " + name + "...")
	caser := cases.Title(language.English)
	name = caser.String(name)

	io.SecureRemoveFile(io.GetCurrentDirectory() + "/frontend/src/services/" + name + ".ts")
	logger.Log("Service deleted.")

}

func serviceTemplate(name string) string {

	var tpl bytes.Buffer
	templates.SERVICE.Execute(&tpl, map[string]interface{}{
		"name": name,
	})

	return tpl.String()
}
