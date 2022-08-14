import { Kafka as KClient, Partitioners } from "kafkajs";

export default class Kafka {
    constructor(configurations) {
        this.configurations = configurations;
        this.kafka = new KClient(this.configurations.cluster)
    }

    async produce(topic, key, data) {
        this.producer = this.kafka.producer({
            createPartitioner: Partitioners.LegacyPartitioner
        })
        await this.producer.connect()
        await this.producer.send({
            topic: topic,
            messages: [{
                "key": key,
                "value": data
            }]
        })
        await this.producer.disconnect()
    }

    async consume(topics, fromBeginning) {
        this.consumer = this.kafka.consumer({
            groupId: this.configurations.consumer.groupId,
            waitForLeaders: true
        })
        await this.consumer.connect()
        await this.consumer.subscribe({ topics: [...topics], fromBeginning: fromBeginning })
        // Find a Better Usage
        await this.consumer.run({
            eachMessage: ({ message }) => {
                console.log(`received message: ${message.value}`)
            },
        })
    }
}

// Usage - Consumer
// import Kafka from "./helpers.js";

// const conf = {
//     cluster: {
//         brokers: [`localhost:9092`],
//         clientId: 'example-producer',
//     },
//     consumer: {
//         groupId: "1234tth"
//     }
// }

// const k1 = new Kafka(conf)
// k1.produce('audit', "87690", '{ "test": "passed" }')
// // k.consume('audit', true)

// Producer

//import Kafka from "./kafka.js";

// const conf = {
//     cluster: {
//         brokers: [`localhost:9092`],
//         clientId: 'example-cons',
//     },
//     consumer: {
//         groupId: "1234tthh"
//     }
// }

// const k = new Kafka(conf)
// // k.produce('audit', "876", '{ "test": "passed" }')
// console.log(k.consume(['audit'], true))