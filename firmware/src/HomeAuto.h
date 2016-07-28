#include <Arduino.h>
#include <PubSubClient.h>
#include <map>

typedef int (*handler)();
class HomeAuto {
private:
  PubSubClient* _client; // The mqtt client
public:
  HomeAuto();
  HomeAuto(PubSubClient& client);
  HomeAuto& setClient(PubSubClient& client);
  void addHandler(const char* name, handler f);
  void callHandler(const char* name);
  void handle();
  void reconnect();
  void msgCallback(char* topic, byte* payload, unsigned int length);
};
