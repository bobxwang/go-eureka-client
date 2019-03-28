Go Eureka Client
================

Based on code from https://github.com/bryanstephens/go-eureka-client .

## Getting started

```go
import (
  "github.com/ArthurHlt/go-eureka-client/eureka"
)

var instanceId string
var appname = "service-go"
var client *eureka.Client

func main() {
  client = eureka.NewClient([]string{
  "http://127.0.0.1:8761/eureka", //From a spring boot based eureka server
    // add others servers here
  })
  hostname := "test.com"
  ip := "127.0.0.1"
  port := 80
  instance := eureka.NewInstanceInfo(hostname, appname, ip, port, 30, false) //Create a new instance to register
  instance.Metadata = &eureka.MetaData{
    Map: make(map[string]string),
  }
  instance.Metadata.Map["foo"] = "bar" //add metadata for example
  client.RegisterInstance(appname, instance) // Register new instance in your eureka(s)
  instanceId = instance.InstanceId
  applications, _ := client.GetApplications() // Retrieves all applications from eureka server(s)
  client.GetApplication(instance.App) // retrieve the application appname
  client.GetInstance(instance.App, instance.InstanceId) // retrieve the instance from "test.com" inside "test"" app
  client.SendHeartbeat(instance.App, instance.InstanceId, 30) // say to eureka that your app is alive (here you must send heartbeat before 30 sec)
}

func exec() {
  client.UnregisterInstance(appname, instancdId)
}
```

**Note:**

- `appname` here is the name of the app
- `instanceId` is the instanceId you register to eureka of the app, the value is composed from ip, port, appname,default value is:  **ip​:appname:port​**
- When calling `RegisterInstance` the `appId` is needed but not used by eureka, this is not the appId but a whatever value, **but we recommend it using appname**,if we using sc gateway to route, it will a path in a url
- SendHeartbeat will using a time.Tick to send a heart to eureka server every 30 second
- When u app exit, u should call **UnregisterInstance** method to tell eureka u app will not serve again

All these strange behaviour come from Eureka.

## Create Client from a config file

You can create from a json file with this form (here we called it `config.json`):

```json
{
  "config": {
    "certFile": "",
    "keyFile": "",
    "caCertFiles": null,
    "timeout": 1000000000,
    "consistency": ""
  },
  "cluster": {
    "leader": "http://127.0.0.1:8761/eureka",
    "machines": [
      "http://127.0.0.1:8761/eureka"
    ]
  }
}
```

And to load it:

```go
client := NewClientFromFile("config.json")
```
