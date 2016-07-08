module.exports = function(app, mqServer) {
app.get('/send/:topic/:payload', function(req,res){
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
});
app.get('/suyash/led/:status', function(req,res){
  var message = {
    topic: 'suyash',
    payload: (req.params.status==="on") ? "LED_ON" : "LED_OFF",
    qos: 0,
    retain: false
  }
  if (!(req.params.status == "on" && req.params.status=="off")){
    res.send("ERR");
  }
  else{
    mqServer.publish(message);
    res.send("Done");
  }

}
