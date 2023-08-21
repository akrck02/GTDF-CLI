package templates

const VIEW = `
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
const VIEW_CORE = `
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
