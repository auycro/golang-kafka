## kafka

kafka topic

```sh
kafka-topics.sh --bootstrap-server=localhost:9092 --list
kafka-topics.sh --bootstrap-server=localhost:9092 --topic=sample_topic --create
kafka-topics.sh --bootstrap-server=localhost:9092 --describe
```

kafka consumer

```sh
kafka-console-consumer.sh --bootstrap-server=localhost:9092 --topic=sample_topic
```

kafka producer

```sh
kafka-console-producer.sh --bootstrap-server=localhost:9092 --topic=sample_topic
```

## mysql

```sh
mysql -u {user} -p
```