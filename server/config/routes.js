/*
routes.js
Route table for the server

@author Suyash Kumar <suyashkumar2003@gmail.com>
*/
var User = require('../models/user');
module.exports = function(app, mqServer, myDeviceEventRouter, jwt) { 
	var deviceMessaging = require('../routes/device-messaging')(mqServer, myDeviceEventRouter);
	var authRoutes = require('../routes/auth.js');
  	// Route Table
	app.get('/sendOnly/:topic/:payload', deviceMessaging.sendMessage);
	app.get('/send/:topic/:payload', deviceMessaging.sendMessageWithResponse); 

	app.get('/test', needsAuth, function(req, res) {
		res.send('Hi ');
  	});

	app.post('/auth', function(req, res) {
		User.findOne({
			email: req.body.email
		}, function(err, user){
			if(err) throw err;
			if(!user) {
				res.json({success:false});
			} else if (user) {
				if(!user.validPassword(req.body.password)) {
					res.json({success:false});
				} else {
					var token = jwt.sign(user, app.get('superSecret'), {expiresIn: "10h"});
					res.json({success: true, token:token});
				}
			}
		});

  	});

	app.get('/users', needsAuth, function(req, res) {
		User.find({}, function(err, users) {
			res.json(users);
		});

	});

	app.get('/setup', function(req, res) {
		var sk = new User({
			email:'sk@test.com' 
		});
		sk.password = sk.generateHash('sk123');
		sk.save(function(err) {
			if(err) throw err;
			res.json({success: true});
		}); 
	});

	function needsAuth(req, res, next){
		
	  // check header or url parameters or post parameters for token
	  var token = req.body.token || req.query.token || req.headers['x-access-token'];

	  // decode token
	  if (token) {

		// verifies secret and checks exp
		jwt.verify(token, app.get('superSecret'), function(err, decoded) {      
		  if (err) {
			return res.json({ success: false, message: 'Failed to authenticate token.' });    
		  } else {
			// if everything is good, save to request for use in other routes
			req.decoded = decoded;    
			next();
		  }
		});

	  } else {

		// if there is no token
		// return an error
		return res.status(403).send({ 
			success: false, 
			message: 'No token provided.' 
		});
		
	  }

	}

	app.use(function(req, res, next) {
	   res.header("Access-Control-Allow-Origin", "http://localhost:3000");
		res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept"); 
		next();
	});
}
function isLoggedIn(req, res, next) {
	if (req.isAuthenticated())
		return next();
	res.status(401);
	res.send('Unauthorized');

}
