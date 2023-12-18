const kafka = require("./client");

async function Produce() {
    const producer = kafka.producer();

    try {
        console.log("Connecting Producer");
        await producer.connect();
        console.log("Producer Connected");

        await producer.send({
            topic: "football-score-updates",
            messages: [
                {
                    key: "[Scored]",
                    value: JSON.stringify({
                        scorerBranch: "IT",
                        rival: "CSE",
                        time: "50:31",
                        scorer: "Tulya Nayak"
                    }),
                    partition: 0,
                }
            ]
        })
        console.log("Message Produced!");

    } catch (error) {
        console.error(error.message);
    }
    console.log("Disconnecting Admin...");
    await producer.disconnect();
}

Produce();