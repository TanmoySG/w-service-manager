import Ajv from "ajv";
import * as filesystem from 'fs';

export default class JSONSchema {
    constructor(filepath) {
        // Load from File
        if (filepath !== undefined) {
            this.schema = JSON.parse(
                filesystem.readFileSync(filepath)
            )
            delete this.schema["$schema"]
            delete this.schema["$id"]
            delete this.schema["title"]
            delete this.schema["description"]
        }
    }

    load(schema) {
        // Load from JSON Object
        this.schema = schema
    }

    validate(data) {
        const ajv = new Ajv()
        this.validator = ajv.compile(this.schema);
        return this.validator(data)
    }
}

// const jsc = new JSONSchema("/Users/tanmoysg/Work/Projects/wunder/w-service-manager/schema/service-onboarding/request.intake.schema.json")
// // jsc.load({
// //     type: "object",
// //     properties: {
// //         foo: { type: "integer" },
// //         bar: { type: "string" }
// //     },
// //     required: ["foo"],
// //     additionalProperties: false,
// // })

// console.log(jsc.validate({
//     foo: 1,
//     bar: "abc"
// }))