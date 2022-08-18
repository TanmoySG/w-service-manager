import { v4 as uuidv4 } from 'uuid';
import JSONSchema from './helpers/schema.js';
import WDB from './helpers/wdb.js';

export default class Request {
    constructor(configuration) {
        this.config = configuration

        this.schema = this.config.schema
        this.jsonSchemaClient = new JSONSchema(this.schema)

        this.dbconfig = this.config.dbconfig
        this.dbClient = new WDB(this.dbconfig)
    }

    new(data, uuid) {
        if (this.jsonSchemaClient.validate(data)) {
            const request = {
                requestID: uuid,
                timestamp: Date.now(),
                status: 'queued',
                serviceName: data.service_name,
                serviceOwners: data.service_owners,
                requestType: data.request_type
            }
            this.dbClient.addData(this.dbconfig.collection, request, function (resp) {
                console.log(resp)
            })
        } else {
            console.log(false)
        }
    }

    id() {
        return `request-${uuidv4()}`
    }
}

// const req = new Request("/Users/tanmoysg/Work/Projects/wunder/w-service-manager/schema/service-onboarding/contract.intake.schema.json",{})

// const samda = {
//     "kind": "request.intake.service-onboarding",
//     "service_name": "fgt",
//     "request_type": "create",
//     "service_owners": ["t", "g", "y"]
// }


// req.new(samda, req.id())