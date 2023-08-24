package templates

import "html/template"

const viewTpl = `
import { ViewUI } from "../../lib/gtdf/views/ViewUI.js";

export default class {{.name}} extends ViewUI {

	private static readonly ID = "{{.name}}";
			
	public constructor() {
		super({
			type: "view",
			id: {{.name}}.ID,
		});
	}

	public show(params : string[]): void {
		console.log("{{.name}} view is showing");
		console.log("Params: ", params);
	}
}	
`
const viewCoreTpl = `
import { ViewCore } from "../../lib/gtdf/views/ViewCore.js";

export default class {{.name}} extends ViewCore {

	public static ping() {
		alert({
			title: "Pong!",
			message: "{{.name}} core is working"
		});
	}

}
`

var VIEW_TEMPLATE = template.Must(template.New("view").Parse(viewTpl))
var VIEW_CORE_TEMPLATE = template.Must(template.New("viewCore").Parse(viewCoreTpl))
