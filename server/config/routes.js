/*
routes.js
Route table for the server

@author Suyash Kumar <suyashkumar2003@gmail.com>
*/
module.exports = function(app, mqServer, myDeviceEventRouter, passport) {
	console.log("test", myDeviceEventRouter.register);
  var deviceMessaging = require('../routes/device-messaging')(mqServer, myDeviceEventRouter);
  var authRoutes = require('../routes/auth.js');
  // Route Table
  app.get('/sendOnly/:topic/:payload', deviceMessaging.sendMessage);
  app.get('/send/:topic/:payload', isLoggedIn, deviceMessaging.sendMessageWithResponse);
  app.get('/logout', authRoutes.logout);
  app.post('/signup', passport.authenticate('local-signup', {
	successRedirect : '/',
	failureRedirect: '/signup'
  }));

  app.post('/login', passport.authenticate('local-login', {
	successRedirect: '/' 
  }));

  app.get('/test', isLoggedIn, function(req, res) {
	res.send('Hi ' + req.user.email);
  });
}

function isLoggedIn(req, res, next) {
	if (req.isAuthenticated())
		return next();
	res.status(401);
	res.send('Unauthorized');

}
