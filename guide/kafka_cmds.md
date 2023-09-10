# Enter kafka container
```shell
docker exec -it <container-id> bash
cd /
```

## Create a topic
```shell
bin/kafka-topics --bootstrap-server <kafka-broker-host> --topic <topic-name> --create
```

e.g.
```shell
bin/kafka-topics --bootstrap-server 127.0.0.1:9092 --topic first_topic --create
```

# Producer
## Produce a message
```shell
bin/kafka-console-producer --bootstrap-server <kafka-broker-host> --topic <topic-name>
```

e.g.
```shell
bin/kafka-console-producer --bootstrap-server localhost:9092 --topic first_topic
# Input a message in shell
```


# Consumer
## Produce a message
```shell
bin/kafka-console-consumer --bootstrap-server <kafka-broker-host> --topic <topic-name>
```

e.g.
```shell
bin/kafka-console-consumer --bootstrap-server localhost:9092 --topic first_topic
```


