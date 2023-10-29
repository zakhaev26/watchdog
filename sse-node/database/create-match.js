const mongoose = require("mongoose")
const FootballScoreModel = require("./FootballSchema")

module.exports = async function CreateMatch(teamA,teamB,scoreA,scoreB) {

    try {
        console.log("Creating!")
        const match = new FootballScoreModel({
            teamA,
            teamB,
            scoreA,
            scoreB
        })
        
        const res = await  match.save();
        console.log("Saved To DB : ",res);

    }catch(e) {
        console.log(e.message)
    }   

}