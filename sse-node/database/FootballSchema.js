const mongoose = require("mongoose");

const FootballScoreSchema = new mongoose.Schema({
    teamA : String,
    teamB : String,
    scoreA : Number,
    scoreB : Number,

})

const FootballScoreModel  = mongoose.model("footballscoreschema",FootballScoreSchema);

module.exports  =  FootballScoreModel;