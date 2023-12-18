const { Kafka } = require("kafkajs")

module.exports = kafka = new Kafka({
    brokers: [`192.168.1.7:9092`],
    clientId: "gc"
})
