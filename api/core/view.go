package core

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/akrck02/GTDF-CLI/api/io"
	"github.com/akrck02/GTDF-CLI/api/logger"
)

func NewView(name string) {

	logger.Log("Current directory: " + io.GetCurrentDirectory())
	if !IsProject(io.GetCurrentDirectory()) {
		logger.Error("You are not in a GTDF project directory.")
		return
	}

	logger.Log("Generating view " + name + "...")

	io.SecureMkdir(io.GetCurrentDirectory()+"/src/views/"+name, 0755)

	template := viewTemplate(strings.Title(name))
	logger.Log(template)

	io.WriteToFile(io.GetCurrentDirectory()+"/src/views/"+name+"/"+strings.Title(name)+".ts", template)

	logger.Log("View generated.")
}

func viewTemplate(name string) string {
	sentence := `

import {UIComponent} from "../../lib/gtd-ts/web/ui-component.js";	

export default class {{.name}} extends UIComponent {
			
	public constructor() {
		super();
	}

	public show(params : string[]): void {
		console.log(" {{.name}} view is showing");
		console.log("Params: ", params);
	}
}	
`

	templ := template.Must(template.New("view").Parse(sentence))
	var tpl bytes.Buffer

	templ.Execute(&tpl, map[string]interface{}{
		"name": name,
	})

	return tpl.String()
}
