#include <Arduino.h>
#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <ESP8266WebServer.h>
#include <ESP8266mDNS.h>
#include <wifi_info.h> // comment this out and fill in the below two lines
#include <PubSubClient.h>

//const char* ssid = "mywifi";
//const char* password = "";

MDNSResponder mdns;
ESP8266WebServer server(80);
WiFiClient client;
PubSubClient pClient(client);
const char* mqtt_server = "10.0.0.98";

void handleRoot() {
  server.send(200, "text/plain", "hello from esp8266!");
}

void handleNotFound(){

  String message = "File Not Found\n\n";
  message += "URI: ";
  message += server.uri();
  message += "\nMethod: ";
  message += (server.method() == HTTP_GET)?"GET":"POST";
  message += "\nArguments: ";
  message += server.args();
  message += "\n";
  for (uint8_t i=0; i<server.args(); i++){
    message += " " + server.argName(i) + ": " + server.arg(i) + "\n";
  }
  server.send(404, "text/plain", message);

}
void reconnect() {
  // Loop until we're reconnected
  while (!pClient.connected()) {
    Serial.print("Attempting MQTT connection...");
    // Attempt to connect
    if (pClient.connect("ESP8266Client")) {
      Serial.println("connected");
      // Once connected, publish an announcement...
      pClient.publish("outTopic", "hello world");
      // ... and resubscribe
      pClient.subscribe("inTopic");
      pClient.subscribe("/lights");
    } else {
      Serial.print("failed, rc=");
      Serial.print(pClient.state());
      Serial.println(" try again in 5 seconds");
      // Wait 5 seconds before retrying
      delay(5000);
    }
  }
}
void callback(char* topic, byte* payload, unsigned int length) {
  Serial.print("Message arrived [");
  Serial.print(topic);
  Serial.print("] ");
  for (int i = 0; i < length; i++) {
    Serial.print((char)payload[i]);
  }
  Serial.println();
}


void setup(void){

  Serial.begin(115200);

  WiFi.begin(ssid, password);
  Serial.println("");


  // Wait for connection
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.print("Connected to ");
  Serial.println(ssid);
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());

  if (mdns.begin("esp8266", WiFi.localIP())) {
    Serial.println("MDNS responder started");
  }
  // Server Routes:
  server.on("/", handleRoot);

  server.on("/inline", [](){
    server.send(200, "text/plain", "this works as well");
  });
  server.on("/redir", [](){
    server.sendHeader("Location", "http://suyashkumar.com/#"+WiFi.localIP().toString());
    server.send(302, "text/plain", "Location: http://suyashkumar.com/#"+WiFi.localIP());
  });
  server.on("/on", [](){
    digitalWrite(D2, HIGH);
    server.send(200,"text/plain","LED ON");
  });
  server.on("/off",[](){
    digitalWrite(D2,LOW);
    server.send(200,"text/plain","LED OFF");
  });

  server.onNotFound(handleNotFound);

  server.begin();
  Serial.println("HTTP server started");

  pClient.setServer(mqtt_server, 1883);
  pClient.setCallback(callback);
}


void loop(void){
  server.handleClient();
  if (!pClient.connected()) {
    reconnect();
  }
  pClient.loop();

}
