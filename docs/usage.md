# unite [![GoDoc](https://godoc.org/github.com/unit-io/unite?status.svg)](https://godoc.org/github.com/unit-io/unite)

## unite is an open source messaging broker for IoT and other real-time messaging service. unite messaging API is built for speed and security.

unite is a real-time messaging service for IoT connected devices, it is based on MQTT protocol. unite is blazing fast and secure messaging infrastructure and APIs for IoT, gaming, apps and real-time web.

unite can be used for online gaming and mobile apps as it satisfy the requirements for low latency and binary messaging. unite is perfect for the internet of things and internet connected devices.

## unite Clients
- [unite-go](https://github.com/unit-io/unite-go) Go client to pubsub messages over protobuf using GRPC application
- [unite-ws](https://github.com/unit-io/unite-ws) Javascript client to pubsub messages over websocket using MQTT protocol. 

## Quick Start
To build unite from source code use go get command and copy [unite.conf](https://github.com/unit-io/unite/tree/master/unite.conf) to the path unite binary is placed.

> go get -u github.com/unit-io/unite && unite

## Usage
The examples folder has various examples for unite usage. Code snippet is given below to use unite messaging broker with web socket and javascript client.

First you need to register a client Id. To get new client id send connect request using blank client id.

Add below mqtt client script to the web page

>  <script src="https://cdnjs.cloudflare.com/ajax/libs/paho-mqtt/1.0.1/mqttws31.js" type="text/javascript"></script>
>  <script src="../unite.js" type="text/javascript"></script>

Send empty client Id (as shown in the below code snippet) in order to register a new client.

> <script type="text/javascript">client = new Paho.MQTT.Client("127.0.01", 6060, "");</script>
```

    // Initialize new Paho client connection
    client = new Paho.MQTT.Client("127.0.01", 6060, "");

    // Set callback handlers
    client.onConnectionLost = onConnectionLost;
    client.onMessageArrived = onMessageArrived;

    // Connect the client, if successful, call onConnect function
    client.connect({
        onSuccess: onConnect,
    });

```

Next step: connect to unite server using primary client Id and send request to generate secondary client Id. The code snippet is given below to request client Id. Client id request type is 0, as you cannot request primary client Id of type 1. You can send multiple requests to generate secondary client Ids.

Note, for topic isolation use client Ids generated from distinct primary client Ids.

```

    client = new Paho.MQTT.Client("127.0.01", 6060, "<<primary clientid>>");
    payload = JSON.stringify({"type":"0"});
    message = new Paho.MQTT.Message(payload);
    message.destinationName = "unite/clientid";
    client.send(message);

```

To subscribe to topic and publish messages to a topic generate key for the topic.

```
    // Once a connection has been made, send payload to request key to publish or subscribe a topic. Pass "rw" to the Type field to set read/write permissions for key on topic.
    payload = JSON.stringify({"topic":"teams.alpha.ch1.u1","type":"rw"});
    message = new Paho.MQTT.Message(payload);
    message.destinationName = "unite/keygen";
    client.send(message);

```

To publish and subscribe to the topic use a valid key.  Key is separated from topic using "/" character.

```
    // Once a connection has been made, publish and subscribe to a topic and use a valid key. Topics are separated by "." character, use * as wildcard character.
    // Subscribe to team alpha channel1 receiver1.
    client.subscribe("<<key>>/teams.alpha.ch1.u1");
    // Publish message to team alpha channel1 receiver1.
    message = new Paho.MQTT.Message("Message for teal alpha channel1 receiver1!");
    message.destinationName = "<<key>>/teams.alpha.ch1.u1";
    client.send(message);

```

Use dot '.' character as topic separator and use three dots '`...`' at the end to subscribe to all topics following the path or use '`*`' character to subscribe to single wildcard topic.

Following are valid topic subscriptions:
Subscribe to team alpha all channels and all receivers
- "key/teams.alpha`...`"

Subscribe to any team channel1 receiver1
- "key/teams.`*`.ch1.u1"

Send messages to channel and channel receivers:

```
    // Once a connection has been made, publish and subscribe to a topic and use a valid key.
    message = new Paho.MQTT.Message("Hello team alpha channel1 receiver1!");
    message.destinationName = "<<key>>/teams.alpha.ch1.u1;
    client.send(message);

```

## Contributing
If you'd like to contribute, please fork the repository and use a feature branch. Pull requests are welcome.

## Licensing
Copyright (c) 2016-2020 Saffat IT Solutions Pvt Ltd. This project is licensed under [MIT License](https://github.com/unit-io/unite/blob/master/LICENSE).
