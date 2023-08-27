package core

import (
	"bytes"
	"strings"

	"github.com/akrck02/GTDF-CLI/api/io"
	"github.com/akrck02/GTDF-CLI/api/logger"
	"github.com/akrck02/GTDF-CLI/api/templates"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func NewView(name string) {

	logger.Log("Generating view " + name + "...")
	caser := cases.Title(language.English)
	name = caser.String(name)

	io.SecureMkdir(io.GetCurrentDirectory()+"/frontend/src/views/"+strings.ToLower(name), 0755)

	template := viewTemplate(name)
	coreTemplate := viewCoreTemplate(name)

	io.WriteFileWith(io.GetCurrentDirectory()+"/frontend/src/views/"+strings.ToLower(name)+"/"+name+".ts", template)
	io.WriteFileWith(io.GetCurrentDirectory()+"/frontend/src/views/"+strings.ToLower(name)+"/"+name+".core.ts", coreTemplate)

	logger.Log("View generated.")
}

func DeleteView(name string) {

	logger.Log("Deleting view " + name + "...")
	caser := cases.Title(language.English)
	name = caser.String(name)

	io.SecureRemoveFile(io.GetCurrentDirectory() + "/frontend/src/views/" + strings.ToLower(name) + "/" + name + ".ts")
	io.SecureRemoveFile(io.GetCurrentDirectory() + "/frontend/src/views/" + strings.ToLower(name) + "/" + name + ".core.ts")
	io.SecureRemoveFile(io.GetCurrentDirectory() + "/frontend/src/views/" + strings.ToLower(name) + "/")

	logger.Log("View deleted.")
}

func viewTemplate(name string) string {

	var tpl bytes.Buffer
	templates.VIEW_TEMPLATE.Execute(&tpl, map[string]interface{}{
		"name": name,
	})

	return tpl.String()
}

func viewCoreTemplate(name string) string {

	var tpl bytes.Buffer
	templates.VIEW_CORE_TEMPLATE.Execute(&tpl, map[string]interface{}{
		"name": name,
	})

	return tpl.String()

}
