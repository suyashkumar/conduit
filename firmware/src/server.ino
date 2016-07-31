#include <Arduino.h>
#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <wifi_info.h> // comment this out and fill in the below two lines
#include <PubSubClient.h>
#include <HomeAuto.h>

#define LED 4

// Fill out the below Github peeps:
//const char* ssid = "mywifi";
//const char* password = "";

WiFiClient client;
PubSubClient pClient(client);
HomeAuto homeAuto("suyash", "10.0.0.98"); // or "suyash", "home.suyash.io"
int ledStatus = 0;

int ledToggle(){
  digitalWrite(LED, (ledStatus) ? LOW : HIGH);
  ledStatus = (ledStatus) ? 0 : 1;
  Serial.println("Toggled");
}

void startWIFI(){
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

}

void setup(void){
  Serial.begin(115200); // Start serial
  pinMode(LED, OUTPUT); // Set LED pin to output

  startWIFI(); // Config/start wifi

  // HomeAuto bindings
  homeAuto.addHandler("ledToggle", &ledToggle);
  homeAuto.setClient(pClient);

}

void loop(void){
  homeAuto.handle();
}
