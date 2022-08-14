import Kafka from "./helpers.js";

const conf = {
    cluster: {
        brokers: [`localhost:9092`],
        clientId: 'example-producer',
    },
    consumer: {
        groupId: "1234tth"
    }
}

const k1 = new Kafka(conf)
k1.produce('audit', "87690", '{ "test": "passed" }')
// k.consume('audit', true)