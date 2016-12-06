# conduit

Conduit allows you to quickly build cloud connected hardware that you can control using a secure RESTful API. The current firmware library runs on WIFI-enabled ESP8266 and Arduino compatible microcontrollers. 

In a nutshell, conduit allows you to directly call arbitrary C functions on your firmware from anywhere in the world via a RESTful API. 

### Sample Project
[smart-lights](https://github.com/suyashkumar/smart-lights) is a sample project that uses this library to switch lights from the cloud. 
![](https://github.com/suyashkumar/smart-lights/blob/master/img/lightswitch.gif)

Currently all firmware<-->server communication has been validated (see project above), but you currently must install an instance of the server yourself (I'll have an instance that takes new accounts up shortly :)). Contact me directly if you're interested or need help!

### Example (In progress)
The basic functionality of this library is straightforward. Start with the provided platformio firmware template and just do the following: 
  
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

  3. [THIS WILL NOT WORK WITHOUT AN ACCOUNT or unless you're running the server locally] Go to `http://home.suyash.io/send/:deviceName/:functionString` and your function will run and any publish messages will be returned in the request response. Note, you must have a valid json web token in a `x-access-token` header. The :deviceName is set when initializing the HomeAuto object (line 27).
