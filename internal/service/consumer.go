package service

import (
    "log"
    "time"
    "github.com/Shopify/sarama"
    "github.com/wvanbergen/kafka/consumergroup"
)

func initConsumer(cgroup string, topic string, zookeeperAddr string)(*consumergroup.ConsumerGroup, error) {
    // consumer config
    config := consumergroup.NewConfig()
    config.Offsets.Initial = sarama.OffsetOldest
    config.Offsets.ProcessingTimeout = 10 * time.Second

    // join to consumer group
    cg, err := consumergroup.JoinConsumerGroup(cgroup, []string{topic}, []string{zookeeperAddr}, config)
    log.Println("ZZOO ADDR", zookeeperAddr)
    if err != nil {
        log.Fatal("!!!!!!!!!!!!!!!!!!!!!!!!!!", err)
        return nil, err
    }

    return cg, err
}

func consume(topic string, cg *consumergroup.ConsumerGroup) {
    for {
        select {
        case msg := <-cg.Messages():
            // messages coming through chanel
            // only take messages from subscribed topic
	    if msg.Topic != topic {
                continue
            }

            log.Println("Topic: ", msg.Topic)
            log.Println("Value: ", string(msg.Value))
	    //Process message ?

            // commit to zookeeper that message is read
            // this prevent read message multiple times after restart
            err := cg.CommitUpto(msg)
            if err != nil {
                log.Println("Error commit zookeeper: ", err.Error())
            }
        }
    }
}
