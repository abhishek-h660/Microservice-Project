const express = require("express")
app = express()
//app.use(express.json)
const bodyParser = require('body-parser')
const {MongoClient} = require("mongodb");
const env = require("dotenv");
const { json } = require("express");
env.config();
//connection uri
const uri = process.env.URI


app.use(
    bodyParser.urlencoded({
      extended: true,
    })
)

app.use(bodyParser.json())
const client = new MongoClient(uri)

function run() {
    client.connect().then(() => {
        try {
            client.db("user").command({ping:1}).then(() => {
                console.log("Connected successfully to mongodb");
            });
        } catch (error) {
            console.log(error)
        }
       
    }).catch(console.error)
    
}

run();

app.get('/', (req, res)=> {
    res.send("Hello Welcome Home");
});

app.post('/signup', (req, res) => {
    email = req.body.email
    password = req.body.password
    const db = client.db("auth-service");
    const userCollection = db.collection("users");

    userCollection.insertOne(req.body).then((result) => {
        console.log(result);
        res.send(result)
    });
});

app.post('/signin', (req, res) => {
    email = req.body.email
    password = req.body.password
    const db = client.db("auth-service");
    const userCollection = db.collection("users");

    userCollection.findOne({"email": email}).then((doc) => {
        if(doc.password == password){
            res.send("Logged In");
        }
    });
});

app.listen(8080, ()=> {
    console.log("app is running on port 8080");
});
