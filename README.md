# home-automation

This is a server and ESP8266/Arduino firmware library that allows for __dead simple__ web-to-device communications. It completley abstracts away network logic needed to remotely call functions on devices and publish data to a server.

The basic functionality of this library is straightforward. Start with the provided platformio template and just do the following: 
  
  1. Write a C function that returns an integer in your Arduino Code:
  
  ```C
  int ledOn(){
    digitalWrite(LED, HIGH);
    homeAuto.publish("LED is now on");
  }
  ```
  2. In your `setup()` function, register your function with the service: 
  
  ```C
  homeAuto.addHandler("ledON", &ledOn);
  ```

  3. Go to `http://home.suyash.io/send/:deviceName/:functionString` and your function will run and any publish messages will be returned in the request response. Note, you must have a valid json web token in a `x-access-token` header. The :deviceName is set when initializing the HomeAuto object (line 27).
