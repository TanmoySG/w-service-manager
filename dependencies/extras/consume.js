import Kafka from "../../onboarding/intake/src/helpers/kafka.js";

const conf = {
    cluster: {
        brokers: [`localhost:9092`],
        clientId: 'example-producer',
    },
    consumer: {
        groupId: "intake-oberver"
    }
}

const k1 = new Kafka(conf)
k1.consume(['intake'], true, function(message){
    console.log(message)
})