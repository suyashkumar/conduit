module.exports = function(app, mqServer) {
  var deviceMessaging = require('../routes/device-messaging')(mqServer);
  // Route Table
  app.get('/send/:topic/:payload', deviceMessaging.sendMessage);
  app.get('/devices/:device_name/led/:status', deviceMessaging.callFunc);
  app.get('/devices/list', deviceMessaging.listDevices);
}
