import * as filesystem from 'fs';
import { v4 as uuidv4 } from 'uuid';
import JSONSchema from './helpers/schema.js';
import WDB from './helpers/wdb.js';

function getContractTemplate(filepath) {
    const contractTemplate = JSON.parse(
        filesystem.readFileSync(filepath)
    )
    return contractTemplate
}

export default class Request {
    constructor(configuration) {
        this.config = configuration

        this.responseTemplate = getContractTemplate(this.config.responseTemplate)

        this.schema = this.config.schema
        this.jsonSchemaClient = new JSONSchema(this.schema)

        this.dbconfig = this.config.dbconfig
        this.dbClient = new WDB(this.dbconfig)
    }

    id() {
        return `request-${uuidv4()}`
    }

    new(data, uuid, callback) {
        if (this.jsonSchemaClient.validate(data)) {
            const request = {
                requestID: uuid,
                timestamp: Date.now(),
                status: 'queued',
                serviceName: data.service_name,
                serviceOwners: data.service_owners,
                requestType: data.request_type
            }
            const contractTemplate = this.responseTemplate
            this.dbClient.addData(this.dbconfig.collection, request, function (resp) {
                if(resp.status_code === '1'){
                    contractTemplate["request-details"] = {
                        request_id: request.requestID,
                        request_status: request.status,
                        timestamp: request.timestamp
                    }
                    contractTemplate['contract-template']['request_id'] = request.requestID
                    contractTemplate['contract-template']['service']['name'] = data.service_name
                    callback(contractTemplate)
                }else{
                    callback({
                        error: "Theres Some Error"
                    })
                }

            })
        } else {
            callback({
                error: 'Invalid Request - Schema Mismatch'
            })
        }
    }

    // Need to Add Email Validation Function
}