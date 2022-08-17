import { v4 as uuidv4 } from 'uuid';
import JSONSchema from './helpers/schema.js';
import isEmail from '@nickgatzos/is-email'

class Contract {
    constructor(schema, dbconfig) {
        this.schema = schema
        this.jschema = new JSONSchema(this.schema)

        this.dbconfig = dbconfig
    }

    new(contract, uuid) {
        // console.log(this.validate(contract))
        if (this.validate(contract) == false) {
            return false
        }
        contract["contractID"] = uuid
        contract["status"] = 'ready'
        return contract
    }

    validate(data) {
        if (this.jschema.validate(data)) {
            data.developer.admin.forEach(mailid => {
                console.log(mailid, isEmail(mailid))
                if (isEmail(mailid) == false) {
                    return false;
                }
            });
            data.developer.contributor.forEach(mailid => {
                console.log(isEmail(mailid))
                if (isEmail(mailid) == false) {
                    return false;
                }
            });
            return true;
        } else {
            return false
        }
    }

    id() {
        return `contract-${uuidv4()}`
    }

}

// const cntr = new Contract('/Users/tanmoysg/Work/Projects/wunder/w-service-manager/schema/service-onboarding/contract.intake.schema.json')

// const smk = {
//     "kind": "contract.intake.service-onboarding",
//     "request_id": "123",
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
//         "admin": ["tanmoysggmailcom", "tsgupta@mail.com"],
//         "contributor": []
//     }
// }

// console.log(cntr.new(smk, cntr.id()))