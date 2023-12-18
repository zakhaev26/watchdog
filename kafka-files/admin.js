const kafka = require("./client");

async function runAdmin() {
    const admin = kafka.admin();

    try {
        console.log("Connecting Admin...");
        await admin.connect();
        console.log("Admin Connection Successful!");

        console.log("creating topic : football-score-updates!");
        await admin.createTopics({
            topics: [{
                topic: "football-score-updates",
                numPartitions: 2,  
            }
            ]
        })
        console.log("created topic : football-score-updates!");
    } catch (e) {
        console.log(e.message);
    }
    console.log("Disconnecting Admin...");
    await admin.disconnect();

}

runAdmin();