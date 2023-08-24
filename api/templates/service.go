package templates

import "html/template"

const serviceTmpl = `
export default class {{.name}} {

    public static ping(){
        return  {
			title: "Pong!",
			message: "{{.name}} service is working"
		};
    }
}
`

var SERVICE = template.Must(template.New("service").Parse(serviceTmpl))
