const express = require("express");
const app = express();
const PORT = 2609
const bodyParser = require("body-parser");
app.use(bodyParser.urlencoded({ extended: true }));
app.use(express.json());
const cors = require("cors");
app.use(cors())

const PFR  = require('./routes/PFR')
const PBR  = require('./routes/PBR')
const CFR  = require('./routes/CFR')
const CBR  = require('./routes/CBR')

app.use("/publish/football",PFR);
app.use("/publish/basketball",PBR);
app.use("/consume/football",CFR);
app.use("/consume/basketball",CBR);

app.get("/", (_, res) => {
    res.status(200).send("Alive")
})

app.listen(PORT, () => {
    console.clear();
    console.log(`:${PORT}`)
})