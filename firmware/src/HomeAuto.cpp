#include "HomeAuto.h"
#include <map>

typedef int (*handler)();
const char* _mqtt_server = "10.0.0.98";
const char* _name = "suyash";
void msgCallback(char* topic, byte* payload, unsigned int length);
void reconnect();
void removeSpace(char* s);

typedef struct node {
  handler f;
  const char* name;
  node *next;
} node_t;

node_t* root;
node_t* current;

HomeAuto::HomeAuto(){
  root = (node_t *)malloc(sizeof(node_t));
  root->next=0;
  root->name="lol";
  current = root;
}
HomeAuto::HomeAuto(PubSubClient& client){
  setClient(client);

}
void HomeAuto::addHandler(const char* name, handler f){
  node *newNode = (node_t*) malloc(sizeof(node_t));
  newNode->f = f;
  newNode->next=0;
  newNode->name=name;
  current->next=newNode;
  current=newNode;
}
void HomeAuto::callHandler(const char* name){
  node_t* currentInSearch = root;
  while(true){
    if (strcmp(name, currentInSearch->name)==0){
      currentInSearch->f(); // Call function
      break;
    }
    if(currentInSearch->next == 0){
      break;
    }
    currentInSearch = currentInSearch->next;
  }
}
HomeAuto& HomeAuto::setClient(PubSubClient& client){
  this->_client = &client;
  client.setServer(_mqtt_server, 1883);
  //client.setCallback(msgCallback);
  client.setCallback([&](char* topic, byte* payload, unsigned int length){
    Serial.print("HI Message arrived [");
    Serial.print(topic);
    Serial.print("] ");
    char payloadStr[length];
    Serial.println(length);
    for (int i = 0; i < length; i++) {
      Serial.print((char)payload[i]);
      payloadStr[i]=(char)payload[i];
    }
    removeSpace(payloadStr);
    Serial.println(payloadStr);
    this->callHandler(payloadStr);
    this->callHandler("hi");
    Serial.println();

  });
  return *this;
}
void HomeAuto::handle(){
  if (!this->_client->connected()){
    this->reconnect();
  }
  this->_client->loop();
}
void removeSpace(char* s)
{
    for (char* s2 = s; *s2; ++s2) {
        if (*s2 != ' ')
            *s++ = *s2;
    }
    *s = 0;
}
void HomeAuto::reconnect() {
  // Loop until we're reconnected
  while (!this->_client->connected()) {
    Serial.print("Attempting MQTT connection...");
    // Attempt to connect
    if (this->_client->connect("ESP8266Client")) {
      Serial.println("connected");
      // Once connected, publish an announcement...
      this->_client->publish("outTopic", "hello world");
      // Suscribe to topics:
      this->_client->subscribe(_name);
      this->_client->subscribe("suyash/status");
    } else {
      Serial.print("failed, rc=");
      Serial.print(this->_client->state());
      Serial.println(" try again in 5 seconds");
      // Wait 5 seconds before retrying
      delay(5000);
    }
  }
}
