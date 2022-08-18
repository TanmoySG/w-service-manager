import { v4 as uuidv4 } from 'uuid';
import JSONSchema from './helpers/schema.js';
import isEmail from '@nickgatzos/is-email'
import Kafka from './helpers/kafka.js';
import WDB from './helpers/wdb.js';

class Contract {
    constructor(configuration) {
        this.config = configuration

        this.schema = this.config.schema
        this.jsonSchemaClient = new JSONSchema(this.schema)

        this.kafkaConfig = this.config.kafkaConfig
        this.kafkaClient = new Kafka(this.kafkaConfig)

        this.dbconfig = this.config.dbconfig
        this.dbClient = new WDB(this.dbconfig)
    }


    id() {
        return `contract-${uuidv4()}`
    }

    new(contract, uuid) {
        if (this.validate(contract) == false) {
            return false
        }
        contract["contractID"] = uuid
        contract["status"] = 'ready'
        return contract
    }

    register(contract) {
        const uuid = contract.contractID
        this.kafkaClient.produce(this.config.kafkaConfig.topic, uuid,JSON.stringify(contract))
        const updatedRequestData = {
            timestamp: Date.now(),
            status: 'proccessed',
        }
        const marker = {
            "Key": "requestID",
            "Value": contract.request_id
        }
        this.dbClient.updateData(this.dbconfig.collection, marker, updatedRequestData, function (resp) {
            console.log(resp)
        })
    }

    validate(data) {
        if (this.jsonSchemaClient.validate(data)) {
            for (var i = 0; i < data.developer.admin.length; i++) {
                const mailid = data.developer.admin[i]
                if (isEmail(mailid) == false) {
                    return false;
                }
            }
            for (var i = 0; i < data.developer.admin.length; i++) {
                const mailid = data.developer.admin[i]
                if (isEmail(mailid) == false) {
                    return false;
                }
            }
            return true;
        } else {
            return false
        }
    }

}


// const config = {}
// const cntr = new Contract(config)

// const smk = {
//     "kind": "contract.intake.service-onboarding",
//     "request_id": "d9f28bbf-f5f4-470a-ab0b-bc7cffbf3fbb",
//     "service": {
//         "name": "placeholders",
//         "respository": "link.to.placeholders",
//         "details": [
//             "A service to store Block data Files",
//             "Block Storage"
//         ]
//     },
//     "data": {
//         "name": {
//             "access": ["read"],
//             "use": "Primary Identifier"
//         },
//         "service_access_token": {
//             "access": ["read"],
//             "use": "Access Propagation"
//         }
//     },
//     "developer": {
//         "admin": ["tanmoysg@gmail.com", "tsgupta@mail.com"],
//         "contributor": []
//     }
// }

// const cntrct = cntr.new(smk, cntr.id())
// cntr.register(cntrct)