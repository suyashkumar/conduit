
// Set up Express.js
var express = require('express');
var app = express();
var server = require('http').Server(app);

// Load libraries
var bodyParser = require('body-parser');
var morgan = require('morgan');

// Set up logging
app.use(morgan('dev'));

// Set up body parsing
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: false}));

// Set up static content
app.use(express.static(__dirname + '/node_modules')); // client-side frameworks
app.use(express.static(__dirname + '/public/dist/')); // HTML, CSS

// Connect to Mongodb
require('./config/db')();

// Set up MQTT Home Automation connections
var mqServer = require('./mqtt-start.js');
var DeviceEventRouter = require('./device-event-router.js');
var myDeviceEventRouter = new DeviceEventRouter(mqServer);
console.log(myDeviceEventRouter.register);

// Set up app routes
require('./config/routes')(app, mqServer, myDeviceEventRouter);

if (!module.parent) {
  var port = process.env.PORT || 9000; // 9000 as default
  // On Linux make sure you have root to open port 80
  server.listen(port, function() {
    console.log('Listening on port ' + port);
  });
}

exports = module.exports = app;
