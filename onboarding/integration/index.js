import * as fs from "fs";
import Logsmith from "logsmithjs";
import { v4 as uuidv4 } from 'uuid';
import ConfigParser from "./src/helpers/config.js";
import RSA from "./src/helpers/crypto.js";
import Mailer from "./src/helpers/email.js";
import Kafka from "./src/helpers/kafka.js";
import WDB from "./src/helpers/wdb.js";

function getToken() {
    return Math.random().toString(36).substring(2) + Math.random().toString(36).substring(2)
}

function getEmailHTML(filepath) {
    return fs.readFileSync(filepath)
}

function main(configurations) {

    const lConfig = configurations["logConfig"]
    const kConfig = configurations["kafkaConfig"]
    const wConfig = configurations["dbconfig"]
    const mConfig = configurations["smtpConfig"]
    const eConfig = configurations["emailTemplates"]

    const log = new Logsmith({});
    log.fetchConfigFromFile(lConfig)

    log.INFO("Starting SO-Integration Agent...")

    const k = new Kafka(kConfig)
    k.consume([kConfig["topic"]], false, function (contractAndChecksString) {

        const contractAndChecks = JSON.parse(contractAndChecksString)
        const contract = contractAndChecks["contract"]
        const contractID = contractAndChecks["contract_id"]
        const receiver = contract["developer"]["admin"]
        const serviceUUID = `service-${uuidv4()}`
        const serviceName = contract["service"]["name"]

        log.INFO(`Proccessing Contract with ID ${contractID}`)

        const rsa = new RSA()
        const { publicKey, privateKey } = rsa.getKeys()

        const integrationData = {
            contract: JSON.stringify(contract),
            namespace: serviceName,
            publicKey: publicKey,
            serviceID: serviceUUID,
            token: getToken()
        }

        const attachments = [
            {
                filename: '.pvtKey',
                content: privateKey
            },
            {
                filename: 'contract.json',
                content: JSON.stringify(contract)
            },
            {
                filename: 'validity.json',
                content: JSON.stringify(contractAndChecks)
            }
        ]

        const dbClient = new WDB(wConfig)
        dbClient.addData('dev', integrationData, function (resp) {
            if (resp.status_code === '1') {
                const mail = new Mailer(mConfig)
                mail.mail(receiver, "Service Onboarded", getEmailHTML(eConfig["successfull-onboarding"]), attachments, function (msg) {
                    log.SUCCESS(`Service Integration Successfull for Service with ID ${serviceUUID} . Mail Sent to Admin.`)
                })
            } else {
                log.FAILURE(`Service Integration Failed for Contract with ID ${contractID} . Mail Sent to Admin.`)
            }
        })

    })
}


const configurationPath = './config/secrets/config.secrets.json'
const baseConfig = new ConfigParser(configurationPath).getBase()

main(baseConfig)