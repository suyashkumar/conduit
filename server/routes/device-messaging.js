const RESPONSE_WAIT_TIME = 2000; // in ms
module.exports=function(mqServer){
  return {
    // Send generic messages given a topic or a payload
    sendMessage: function(req, res){
      var message = {
    		topic: req.params.topic,
    		payload: req.params.payload.trim(),
    		qos:0,
    		retain:false
    	}
    	mqServer.publish(message, function(){
        console.log('done');
      });
    	res.send('sent');
    },

    // Send generic message to a device with a payload (specifying a remote function),
    // also wait for a response from the device
    sendMessageWithResponse: function(req, res){
      var message = {
    		topic: req.params.topic,
    		payload: req.params.payload,
    		qos:0,
    		retain:false
    	}
      // Publish message to client (req.params.topic),
      // wait to recieve message back from client and return it back
    	mqServer.publish(message , function(){
          mqServer.on('published', function(packet, client) {
            // If message from client of this request, return that message as response
            if (client && client.id.trim() === req.params.topic.trim()) {
              if(!res.headersSent) res.json(packet.payload.toString());
            }
          });
      });
      // If the device response callback above doesn't fire,
      // this will fire and respond with an error after RESPONSE_WAIT_TIME
      setTimeout(function(){
        console.log(res.headersSent);
        if (!res.headersSent) res.status(504).json("ERROR--no response received");
      }, RESPONSE_WAIT_TIME);
    }

  }
}
