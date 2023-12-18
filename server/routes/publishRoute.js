const express = require('express')
const Router = express.Router();

const Produce = require('../producer');

Router.post('/',async (req, res) => {
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

module.exports  = Router