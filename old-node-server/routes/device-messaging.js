const RESPONSE_WAIT_TIME = 2000; // in ms
module.exports=function(mqServer, myDeviceEventRouter){
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
    	res.json({data: 'sent'});
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
		
		// Function to be called when device responds to this RPC
		var onDevicePublish = function(packet, client){
			// If message from client of this request, return that message as response
			if (client && client.id.trim() === req.params.topic.trim()) {
				if(!res.headersSent) res.json({success: true, data: packet.payload.toString()});
			} 
	  	};

		// Register onDevicePublish to be called when the current device publishes data,
		// then publish the RPC message to the device. 
		myDeviceEventRouter.register(req.params.topic + "/device", onDevicePublish, () => (mqServer.publish(message))); 

      // If the device response callback above (onDevicePublish) doesn't fire,
      // this will fire and respond with an error after RESPONSE_WAIT_TIME
      setTimeout(function(){
        console.log(res.headersSent);
        if (!res.headersSent) res.status(504).json({success: false, data: "ERROR--no response received"});
      }, RESPONSE_WAIT_TIME);
    }

  }
}
