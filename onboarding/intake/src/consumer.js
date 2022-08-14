import Kafka from "./kafka.js";

const conf = {
    cluster: {
        brokers: [`localhost:9092`],
        clientId: 'example-cons',
    },
    consumer: {
        groupId: "1234tthh"
    }
}

const k = new Kafka(conf)
// k.produce('audit', "876", '{ "test": "passed" }')
console.log(k.consume(['audit'], true))