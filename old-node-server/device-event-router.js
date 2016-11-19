class DeviceEventRouter {
	constructor(mqServer) {
		this.mqServer = mqServer;
		this.eventMap = {};
		this.handleEvent = this.handleEvent.bind(this); 

		mqServer.on('published', this.handleEvent);
	}

	/*
	 * This function is called to register a new callback
	 * function (publishCallback) for device publish events 
	 * with the streamString topic. 
	 */
	register(streamString, publishCallback, callback) { 
		this.eventMap[streamString] = publishCallback; 
		callback();
	}
	
	/* 
	 * Called whenever a device publishes data. This function 
	 * ensures that the proper registered callback function in 
	 * in eventMap called for this publish event (based on the 
	 * publish topic)
	 */
	handleEvent(packet, client) { 
		if(client && this.eventMap.hasOwnProperty(packet.topic.trim())) {
			this.eventMap[packet.topic.trim()](packet, client); // Fire registered handler 	
		} else {
			// Do other stuff like add to proper data stream
		} 
	} 
}

module.exports = DeviceEventRouter;
