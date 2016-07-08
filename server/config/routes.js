module.exports = function(app, serverA) {
 app.get('/hello', function(req,res){
 res.send('hello');
 });
app.get('/send', function(req,res){
	var message = {
		topic: 'inTopic',
		payload: 'abc',
		qos:0,
		retain:false
	}
	serverA.publish(message, function(){
    console.log('done');
  });
	res.send('sent');
});
 }
