package templates

const SERVICE = `
export default class {{.name}} {

    public static ping(){
        return  {
			title: "Pong!",
			message: "{{.name}} service is working"
		};
    }
}

`
