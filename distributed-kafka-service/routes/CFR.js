const express = require('express')
const Router = express.Router();
const { Consume, eventEmitter } = require('../consumer');
const __TOPIC__ = 'football';

Router.get("/", async (req, res) => {
    console.log("Aya")
    res.setHeader("Content-Type", "text/event-stream");
    res.write("Ok")

    req.on("close", () => {
        console.log("Client connection closed");
    });

    try {
        await Consume(__TOPIC__,"cgrp2");
        res.write("data: Initial message\n\n");

        eventEmitter.on(__TOPIC__, (message) => {
            res.write(`data: ${message}\n\n`);
        });
    } catch (error) {
        console.error("Error in /subscribe route:", error);
        res.status(500).end();
    }

})

module.exports = Router