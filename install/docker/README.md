## Reference
This doc is copy from https://www.baeldung.com/ops/kafka-docker-setup.

## Overview

In this tutorial, we'll learn how to do an Apache Kafka setup using Docker.

## Single Node Setup

A single node Kafka broker setup would meet most of the local development needs, so let's start by learning this simple setup.

### 2.1. docker-compose.yml Configuration

To start an Apache Kafka server, we'd first need to start a Zookeeper server.

We can configure this dependency in a docker-compose.yml file, which will ensure that the Zookeeper server always starts before the Kafka server and stops after it. The docker-compose file is under `single-node` directory.

In this setup, our Zookeeper server is listening on port=2181 for the kafka service, which is defined within the same container setup. However, for any client running on the host, it'll be exposed on port 22181.

Similarly, the kafka service is exposed to the host applications through port 29092, but it is actually advertised on port 9092 within the container environment configured by the KAFKA_ADVERTISED_LISTENERS property.

### 2.2 Start Kafka Server
Start the Kafka server by the docker-compose command:
```bash
$ docker-compose up -d
Creating network "kafka_default" with the default driver
Creating kafka_zookeeper_1 ... done
Creating kafka_kafka_1     ... done 
```

Use the nc command to verify that both the servers are listening on the respective ports:
```bash
$ nc -z localhost 22181
Connection to localhost port 22181 [tcp/*] succeeded!
$ nc -z localhost 29092
Connection to localhost port 29092 [tcp/*] succeeded!
```

Additionally, we can also check the verbose logs while the containers are starting up and verify that the Kafka server is up:
```bash
$ docker-compose logs kafka | grep -i started
kafka_1      | [2021-04-10 22:57:40,413] DEBUG [ReplicaStateMachine controllerId=1] Started replica state machine with initial state -> HashMap() (kafka.controller.ZkReplicaStateMachine)
kafka_1      | [2021-04-10 22:57:40,418] DEBUG [PartitionStateMachine controllerId=1] Started partition state machine with initial state -> HashMap() (kafka.controller.ZkPartitionStateMachine)
kafka_1      | [2021-04-10 22:57:40,447] INFO [SocketServer brokerId=1] Started data-plane acceptor and processor(s) for endpoint : ListenerName(PLAINTEXT) (kafka.network.SocketServer)
kafka_1      | [2021-04-10 22:57:40,448] INFO [SocketServer brokerId=1] Started socket server acceptors and processors (kafka.network.SocketServer)
kafka_1      | [2021-04-10 22:57:40,458] INFO [KafkaServer id=1] started (kafka.server.KafkaServer)
```

### 2.3. Connection Using Kafka Tool
> https://kafkatool.com/download.html

Finally, let's use the Kafka Tool GUI utility to establish a connection with our newly created Kafka server, and later, we'll visualize this setup:

We must note that we need to use the Bootstrap servers property to connect to the Kafka server listening at port 29092 for the host machine.

### 2.4 Stop Kafka Server
Run
```bash
$ docker-compose stop
```

## 3. Kafka Cluster Setup
For more stable environments, we'll need a resilient setup. Let's extend our docker-compose.yml file to create a multi-node Kafka cluster setup.

### 3.1. docker-compose.yml Configuration
A cluster setup for Apache Kafka needs to have redundancy for both Zookeeper servers and the Kafka servers.

So, let's add configuration for one more node each for Zookeeper and Kafka services. The docker-compose file is under `cluster` folder.

We must ensure that the service names and KAFKA_BROKER_ID are unique across the services.

Moreover, each service must expose a unique port to the host machine. Although zookeeper-1 and zookeeper-2 are listening on port 2181, they're exposing it to the host via ports 22181 and 32181, respectively. The same logic applies for the kafka-1 and kafka-2 services, where they'll be listening on ports 29092 and 39092, respectively.

### 3.2. Start the Kafka Cluster

Start the Kafka Cluster by using the docker-compose command:
```bash
$ docker-compose up -d
Creating network "kafka_default" with the default driver
Creating kafka_zookeeper-1_1 ... done
Creating kafka_zookeeper-2_1 ... done
Creating kafka_kafka-2_1     ... done
Creating kafka_kafka-1_1     ... done
```

Once the cluster is up, let's use the Kafka Tool to connect to the cluster by specifying comma-separated values for the Kafka servers and respective ports:

## 4. Conclusion

In this article, we used the Docker technology to create single node and multi-node setups of Apache Kafka.

We also used the Kafka Tool to connect and visualize the configured broker server details.
