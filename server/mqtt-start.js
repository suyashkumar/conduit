var mosca = require('mosca');
var settings = {
	port: 1883
}
var mqServer = new mosca.Server(settings);

mqServer.on('clientConnected', function(client){
	console.log(client.id);
  var message = {
    topic: '/lights',
    payload: 'abc',
    qos:0,
    retain:false
  }
  mqServer.publish(message, function(){
    console.log('done');
  });
});
mqServer.on('published', function(packet,client){
  console.log("Client:","Published",packet.payload.toString());
  console.log(packet.payload);
  if(client){
    console.log(client.id);
  }
});
module.exports=mqServer; // Export the init'd mqServer object 
