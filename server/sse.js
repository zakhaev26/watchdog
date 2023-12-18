const express = require("express");
const app = express();
const PORT = 2609
const bodyParser = require("body-parser");
app.use(bodyParser.urlencoded({ extended: true }));
app.use(express.json());
const Produce = require("./producer")
const {Consume , eventEmitter} = require("./consumer");
const cors = require("cors");
const publishRoute = require("./routes/publishRoute");
const subscribeRoute = require("./routes/subscribeRoute");
app.use(cors())
app.use("/publish",publishRoute);
app.use("/subscribe",subscribeRoute);


app.get("/", (_, res) => {
    res.status(200).send("Alive")
})

// app.post('/publish',async (req, res) => {
//     console.log(req.body)
//     try{
//         const { event, scorerBranch, rival, time, scorer } = req.body;
//         await Produce( event, scorerBranch, rival, time, scorer );
//         res.status(200).json({
//             "published":"true",
//         })
//     }catch(e) {
//         res.status(500).json({
//             error:`${e}`
//         })
//     }
// })

// app.get("/subscribe", async(req,res) =>{
//     console.log("Aya")
//     res.setHeader("Content-Type", "text/event-stream");
//     res.write("Ok")
    
//     req.on("close", () => {
//         console.log("Client connection closed");
//     });

//     try {
//         await Consume();

//         res.write("data: Initial message\n\n");

//         eventEmitter.on("message", (message) => {
//             res.write(`data: ${message}\n\n`);
//         });
//     } catch (error) {
//         console.error("Error in /subscribe route:", error);
//         res.status(500).end();
//     }

// })

app.listen(PORT, () => {
    console.clear();
    console.log(`:${PORT}`)
})