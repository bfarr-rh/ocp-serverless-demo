var express = require('express');
var exphbs  = require('express-handlebars');
var app = express();
var os = require("os");
var morgan  = require('morgan');
const PORT = process.env.PORT || 3000;
const URL = process.env.URL || "http://happydayz.serverless-demo.svc.cluster.local";


app.engine('handlebars', exphbs({defaultLayout: 'main'}));
app.set('view engine', 'handlebars');
app.use(express.static('static'));
app.use(morgan('combined'));

const fetch = require('node-fetch');
const { log } = require('console');

// Configuration
// var port = process.env.PORT || 8080;
// var message = process.env.MESSAGE || "I 💕 NodeJS!!!";
// var lang = "NodeJS"
var message = "";
var lang = ""


let settings = { method: "Get" };


app.get('/', function (req, res) {


  (async () => {
    try {
      await fetch(URL, settings)
      .then(res => res.json())
      .then((json) => {
          lang = json.app;
          message = json.message;
      });
    } catch (error) {
      if (error.name === 'AbortError') {
        console.log('request was aborted');
      }
      message = "I 💕 NodeJS @ TPG!!!";
      lang = "NodeJS"
    }
  })();
      console.log(message);
      console.log(lang);
      res.render('home', { 
      message: message,
      platform: os.type(),
      release: os.release(),
      hostName: os.hostname(),
      lang: lang,
    });
});

// Set up listener
app.listen(PORT, function () {
  console.log("Listening on: http://%s:%s", os.hostname(), PORT);
});