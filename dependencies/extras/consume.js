import Kafka from "../../onboarding/intake/src/helpers/kafka.js";

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
k1.consume(['intake'], true)