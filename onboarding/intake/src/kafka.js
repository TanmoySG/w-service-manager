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