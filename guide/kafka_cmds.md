# Enter kafka container
```shell
docker exec -it <container-id> bash
cd /
```

## Create a topic
Command:
```shell
bin/kafka-topics --bootstrap-server <kafka-broker-host> --topic <topic-name> --create
```

E.g. create the first topic.
```shell
bin/kafka-topics --bootstrap-server 127.0.0.1:9092 --topic first_topic --create
```

# Producer
Command:
```shell
bin/kafka-console-async_producer --bootstrap-server <kafka-broker-host> --topic <topic-name>
```

E.g. produce some messages.
```shell
bin/kafka-console-async_producer --bootstrap-server localhost:9092 --topic first_topic
```

Then type in some message in the console. 


# Consumer
Command:
```shell
bin/kafka-console-consumer --bootstrap-server <kafka-broker-host> --topic <topic-name>
```

E.g. Consume messages
```shell
bin/kafka-console-consumer --bootstrap-server localhost:9092 --topic first_topic
```

E.g. Consume messages from the very beginning
```shell
bin/kafka-console-consumer --bootstrap-server localhost:9092 --topic first_topic --from-beginning
```

