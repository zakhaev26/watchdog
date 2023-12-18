const express = require('express')
const Router = express.Router();

const Produce = require('../producer');
const __TOPIC__ = 'basketball';
Router.post('/',async (req, res) => {
    console.log(req.body)
    try{
        const { event, scorerBranch, rival, time, scorer } = req.body;
        console.log({ event, scorerBranch, rival, time, scorer})
        await Produce( event, scorerBranch, rival, time, scorer ,__TOPIC__);
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