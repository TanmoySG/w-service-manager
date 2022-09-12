import Kafka from "../../onboarding/intake/src/helpers/kafka.js";

const conf = {
    cluster: {
        brokers: [`localhost:9092`],
        clientId: 'example-c',
    },
    consumer: {
        groupId: "1234ttggh"
    }
}

const k1 = new Kafka(conf)
const tfr = k1.produce('intake', "103f", JSON.stringify({"test":"passs"}))
console.log(tfr)

// k1.consume(['intake'], true, function (message) {
//     console.log(message)
// })