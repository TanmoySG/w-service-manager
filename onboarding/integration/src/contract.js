import isEmail from '@nickgatzos/is-email';
import { v4 as uuidv4 } from 'uuid';

import Kafka from './helpers/kafka.js';
import JSONSchema from './helpers/schema.js';
import WDB from './helpers/wdb.js';

export default class Contract {
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

    register(contract, uuid, callback) {
        if (this.validate(contract) == false) {
            callback({
                error: "Invalid Contract"
            })
        } else {
            const marker = {
                "Key": "requestID",
                "Value": contract.request_id
            }

            // For Use inside Callback
            // https://stackoverflow.com/questions/20279484/how-to-access-the-correct-this-inside-a-callback
            var self = this
            
            this.dbClient.getData(this.dbconfig.collection, marker, function (data, err) {
                if (err !== undefined) {
                    callback({
                        error: "Invalid Request ID"
                    })
                } else {
                    const serviceOwnersInContract = contract.developer.admin
                    if(self.validateOwners(data.serviceOwners, serviceOwnersInContract)) {
                        contract["contractID"] = uuid
                        contract["status"] = 'ready'

                        self.kafkaClient.produce(self.config.kafkaConfig.topic, uuid, JSON.stringify(contract))

                        const updatedRequestData = {
                            timestamp: Date.now(),
                            status: 'proccessed',
                        }

                        self.dbClient.updateData(self.dbconfig.collection, marker, updatedRequestData, function (resp) {
                            if (resp.status_code === '1') {
                                const contractResponse = {
                                    contract_id: uuid,
                                    request_id: contract.request_id,
                                    request_status: updatedRequestData.status,
                                    timestamp: updatedRequestData.timestamp
                                }
                                callback(contractResponse)
                            } else {
                                callback({
                                    error: "There's Some Error with the DB. Hold on tight!"
                                })
                            }
                        })
                    } else {
                        callback({
                            error: "Service Owners in Request do not match."
                        })
                    }

                }
            })
        }
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

    validateOwners(owners, contractOwners) {
        return contractOwners.every(val => owners.includes(val))
    }
}