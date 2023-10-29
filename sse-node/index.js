
const app = require('express')();
const FootballScoreModel = require('./database/FootballSchema');
const connToDB = require("./database/connection");
const createMatch = require('./database/create-match');
app.use(require("express").json());
app.use(cors())
connToDB();

app.get("/", (req, res) => {
    res.end("Hello")
})

const changeStream = FootballScoreModel.watch();

app.get("/stream", (req, res) => {
    res.setHeader("Content-Type", "text/event-stream");

    changeStream.on('change', (change) => {
        console.log("Changed Event : ", change.fullDocument);
        res.json(change);
    })
})

app.post("/post-data", (req, res) => {
    const { teamA ,teamB, scoreA , scoreB } = req.body;
    createMatch(teamA,teamB,scoreA,scoreB);
    res.redirect("/");
})

app.listen(8080, () => {
    console.clear();
    console.log(`Running on port 8080`);
})