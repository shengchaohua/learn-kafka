# start kafka
docker-compose up -d

# stop kafka
docker-compose stop

# run a kafka consumer
docker exec -it single-node-kafka-1 bash

# in container
cd /
bin/kafka-console-consumer --bootstrap-server=127.0.0.1:29092 -topic=first
# consume from beginning
bin/kafka-console-consumer --bootstrap-server=127.0.0.1:29092 -topic=first --from-beginning
