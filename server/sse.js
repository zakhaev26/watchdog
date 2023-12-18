const express = require("express");
const app = express();
const PORT = 2609
const bodyParser = require("body-parser");
app.use(bodyParser.urlencoded({ extended: true }));
app.use(express.json());
const Produce = require("./producer")


app.get("/", (req, res) => {
    res.status(200).send("Alive")
})

app.post('/publish',async (req, res) => {
    console.log(req.body)
    try{
        const { event, scorerBranch, rival, time, scorer } = req.body;
        await Produce( event, scorerBranch, rival, time, scorer );
        res.status(200).json({
            "published":"true",
        })
    }catch(e) {
        res.status(500).json({
            error:`${e}`
        })
    }
})

app.get("/subsribe",(req,res) =>{
    
})

app.listen(PORT, () => {
    console.clear();
    console.log(`:${PORT}`)
})