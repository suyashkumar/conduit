module.exports = function(app, serverA) {
 app.get('/hello', function(req,res){
 res.send('hello');
 });
app.get('/send/:topic/:payload', function(req,res){
	var message = {
		topic: req.params.topic,
		payload: req.params.payload,
		qos:0,
		retain:false
	}
	serverA.publish(message, function(){
    console.log('done');
  });
	res.send('sent');
});
 }
