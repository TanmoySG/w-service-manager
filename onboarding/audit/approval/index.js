import Kafka from "../../intake/src/helpers/kafka.js";

const config = {
    cluster: {
        brokers: [`localhost:9092`],
        clientId: 'approval-audit-listener',
    },
    consumer: {
        groupId: "approval-audit-observer"
    }
}

const k = new Kafka(config)
k.consume(['audit'], false, function (contract) {
    console.log("Retrived Contract", contract["contract_id"])
    k.produce("integrate", contract["contract_id"], contract)
})