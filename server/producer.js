const kafka = require("./client");
const readline = require("readline")

async function Produce(event,scorerBranch,rival,time,scorer) {
    const producer = kafka.producer();

    try {
        console.log("Connecting Producer");
        await producer.connect();
        console.log("Producer Connected");
            await producer.send({
                topic: "football-score-updates",
                messages: [
                    {
                        key: `[${event}]`,
                        value: JSON.stringify({
                            scorerBranch,
                            rival,
                            time,
                            scorer
                        }),
                        partition: 0,
                    }
                ]
            })
    } catch (error) {
        console.error(error.message);
        console.log("Disconnecting Producer...");
        await producer.disconnect();
    }
}

// Produce()
module.exports = Produce;