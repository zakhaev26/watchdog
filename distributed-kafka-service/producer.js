const kafka = require("./client");

async function Produce(event,scorerBranch,rival,time,scorer,topic) {
    const producer = kafka.producer();

    try {
        console.log("Connecting Producer");
        await producer.connect();
        console.log("Producer Connected");
            await producer.send({
                topic,
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
        console.log("Message sent!")
    } catch (error) {
        console.error(error.message);
    }
    console.log("Disconnecting Producer...");
    await producer.disconnect();
}

// Produce()
module.exports = Produce;