const kafka = require("./client");

async function runAdmin(topic) {
    const admin = kafka.admin();

    try {
        console.log("Connecting Admin...");
        await admin.connect();
        console.log("Admin Connection Successful!");

        console.log(`creating topic : ${topic} !`);
        await admin.createTopics({
            topics: [{
                topic,
                numPartitions: 2,  
            }
            ]
        })
        console.log(`created topic : ${topic}!`);
    } catch (e) {
        console.log(e.message);
    }
    console.log("Disconnecting Admin...");
    await admin.disconnect();
}

runAdmin('football');