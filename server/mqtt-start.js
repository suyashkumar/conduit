var mosca = require('mosca');
var settings = {
	port: 1883
};
var mqServer = new mosca.Server(settings); // Init a mosca mqtt server

mqServer.on('clientConnected', function(client) {
	console.log("Client Connected:", client.id);
});

module.exports=mqServer; // Export the init'd mqServer object
