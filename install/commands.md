## Topic
Command:
```shell
bin/kafka-topics [--option]
```

Options:


Examples:
1) create a topic
```shell
bin/kafka-topics --bootstrap-server 127.0.0.1:9092 --topic first_topic --create
```

## Producer
Command:
```shell
bin/kafka-console-async_producer --bootstrap-server <kafka-broker-host> --topic <topic-name>
```

Example:

1) produce some messages.

```shell
bin/kafka-console-async_producer --bootstrap-server localhost:9092 --topic first_topic
```

Then type in some message in the console. 


## Consumer
Command:
```shell
bin/kafka-console-consumer --bootstrap-server <kafka-broker-host> --topic <topic-name>
```

Example:

1) consume messages
```shell
bin/kafka-console-consumer --bootstrap-server localhost:9092 --topic first_topic
```

