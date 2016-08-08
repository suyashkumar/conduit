/*
routes.js
Route table for the server

@author Suyash Kumar <suyashkumar2003@gmail.com>
*/
module.exports = function(app, mqServer, myDeviceEventRouter) {
	console.log("test", myDeviceEventRouter.register);
  var deviceMessaging = require('../routes/device-messaging')(mqServer, myDeviceEventRouter);
  // Route Table
  app.get('/sendOnly/:topic/:payload', deviceMessaging.sendMessage);
  app.get('/send/:topic/:payload', deviceMessaging.sendMessageWithResponse);
}
