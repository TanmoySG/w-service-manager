import { v4 as uuidv4 } from 'uuid';
import JSONSchema from './helpers/schema.js';
import WDB from './helpers/wdb.js';

class Request {
    constructor(schema, dbconfig) {
        this.schema = schema
        this.jschema = new JSONSchema(this.schema)

        this.dbconfig = dbconfig
        this.db = new WDB(this.dbconfig)

    }

    new(data, uuid) {
        if (this.jschema.validate(data)) {
            const request = {
                requestID: uuid,
                timestamp: Date.now(),
                status: 'queued',
                serviceName: data.service_name,
                serviceOwners: data.service_owners,
                requestType: data.request_type
            }
            this.db.addData(this.dbconfig.collection, request, function (resp) {
                console.log(resp)
            })
        } else {
            console.log(false) 
        }
    }

    id() {
        return uuidv4();
    }
}

// const req = new Request()

// const samda = {
//     "kind": "request.intake.service-onboarding",
//     "service_name": "fgt",
//     "request_type": "create",
//     "service_owners": ["t", "g", "y"]
// }


// req.new(samda, req.id())