var deviceMap = {
  'suyash': {
    functions: ['led']
  }
}
module.exports=function(mqServer){
  return {
    // Send generic messages given a topic or a payload
    sendMessage: function(req,res){
      var message = {
    		topic: req.params.topic,
    		payload: req.params.payload,
    		qos:0,
    		retain:false
    	}
    	mqServer.publish(message, function(){
        console.log('done');
      });
    	res.send('sent');
    },
    callFunc: function(req,res){
      var message = {
        topic: req.params.device_name,
        payload: (req.params.status==="on") ? "LED_ON" : "LED_OFF",
        qos: 0,
        retain: false
      }
      if (!(req.params.status == "on" || req.params.status=="off")){
        res.send("ERR");
      }
      else{
        mqServer.publish(message);
        res.send("Done");
      }
    },
    listDevices: function(req,res){
      res.json(deviceMap);

    }
  }
}
