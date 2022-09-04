import Kafka from "../../onboarding/intake/src/helpers/kafka.js";

const conf = {
    cluster: {
        brokers: [`localhost:9092`],
        clientId: 'example-producer',
    },
    consumer: {
        groupId: "sudit-oberver"
    }
}

const k1 = new Kafka(conf)
k1.consume(['audit'], true, function(message){
    console.log(message)
})