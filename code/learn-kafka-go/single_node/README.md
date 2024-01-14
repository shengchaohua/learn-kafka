# Start kafka
Use single-node kafka.

Start kafka broker:
```shell
cd install/single-node
docker-compose up -d
```

Stop kafka broker:
```shell
cd install/single-node
docker-compose stop
```

# Create topic
1) Enter kafka container
```shell
docker ps # get CONTAINER ID/NAME
docker exex -it <CONTAINER ID/NAME> bash
```

2) Find kafka cmds in `/bin/` and create the 1st topic:
```shell
$ cd /
$ bin/kafka-topics --bootstrap-server 127.0.0.1:9092 --topic first_topic --create
```

3) List topics
```shell
$ bin/kafka-topics --bootstrap-server 127.0.0.1:9092 --list
first_topic
```

## Producer
### Async producer:
Start producer:
```shell
learn-kafka-go run async_producer/main.learn-kafka-go
```

### Sync producer
Start producer
```shell
learn-kafka-go run producer/main.learn-kafka-go
```

## Consumer
Start consumer:
```shell
learn-kafka-go run consumer/main.learn-kafka-go
```
