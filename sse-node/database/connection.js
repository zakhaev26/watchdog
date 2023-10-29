const mongoose = require("mongoose");


module.exports = async function connToDB() {
    try {
        
        await mongoose.connect("mongodb+srv://soubhik:0iKvSaUZxIQFSTeU@testcluster.t47d55j.mongodb.net/?retryWrites=true&w=majority");
        console.log("Connected to database!")
    } catch (error) {
        console.log(error)
        exit(1)
    }
}


