const kafka = require("./client");

async function Consume() {
    const consumer = kafka.consumer({ groupId: "consumer-1" });

    try {
        console.log("Consumer connecting..");
        await consumer.connect();
        console.log("Consumer connected! ")

        await consumer.subscribe({
            topics: ["football-score-updates"],
            fromBeginning: true
        })

        await consumer.run({
            eachMessage: async ({ topic, partition, message, heartbeat, pause }) => {
                console.log(`[${topic}]/[PART:${partition}]:${message.value.toString()}`);
            },
        })

    } catch (error) {
        console.log(error);
        console.log("Disconnecting Consumer")

        await consumer.disconnect();

    }
}

Consume();