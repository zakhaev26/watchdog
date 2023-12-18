const kafka = require("./client");
const EventEmitter = require("events")

const eventEmitter = new EventEmitter();

async function Consume(topic,grpID) {
    const consumer = kafka.consumer({ groupId: grpID });

    try {
        console.log("Consumer connecting..");
        await consumer.connect();
        console.log("Consumer connected! ")

        await consumer.subscribe({
            topics: [`${topic}`],
            fromBeginning: true
        })

        console.log("Consumer subscribed! ")

        await consumer.run({
            eachMessage: async ({ topic, partition, message, heartbeat, pause }) => {
                console.log(`[${topic}]/[PART:${partition}]:${message.value.toString()}`);
                eventEmitter.emit(topic, `[${topic}]/[PART:${partition}]:${message.value.toString()}`);
            },
        })

    } catch (error) {
        console.log(error);
        console.log("Disconnecting Consumer")
        await consumer.disconnect();
    }
}

// Consume();
module.exports = { Consume, eventEmitter }