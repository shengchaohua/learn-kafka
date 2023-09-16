# Start kafka
Use single-node kafka.

Start kafka broker:
```shell
cd guide/install-kafka-single-node
docker-compose up -d
```

Stop kafka broker:
```shell
cd guide/install-kafka-single-node
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
go run async_producer/main.go
```

### Sync producer
Start producer
```shell
go run sync_producer/main.go
```

## Consumer
Start consumer:
```shell
go run consumer/main.go
```
