package org.example.kafka.producer;

import org.apache.kafka.clients.producer.*;
import org.apache.kafka.clients.producer.internals.Sender;
import org.apache.kafka.common.serialization.StringSerializer;
import org.example.kafka.KafkaConfig;

import java.util.Properties;

public class CustomProducerCallback {
    public static void main(String[] args) {
        // 1. 创建kafka生产者的配置对象
        Properties properties = new Properties();

        // 2. 给kafka配置对象添加配置信息：bootstrap.servers
        properties.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, KafkaConfig.Host);

        // key,value序列化（必须）：key.serializer，value.serializer
        properties.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());
        properties.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());

        // 3. 创建kafka生产者对象
        KafkaProducer<String, String> kafkaProducer = new KafkaProducer<String, String>(properties);

        // 4. 调用send方法,发送消息
        for (int i = 0; i < 5; i++) {
            kafkaProducer.send(new ProducerRecord<>("first", "value " + i), new Callback() {
                @Override
                public void onCompletion(RecordMetadata metadata, Exception exception) {
                    if (exception == null) {
                        System.out.println("主题：" + metadata.topic() + ", "  + "分区：" + metadata.partition());
                    } else {
                        exception.printStackTrace();
                    }
                }
            });
        }

        // 5. 关闭资源
        kafkaProducer.close();
    }
}