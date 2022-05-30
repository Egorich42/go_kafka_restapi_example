package service

import ("github.com/Egorich42/go_kafka_restapi_example/container"
	"os"
	"github.com/Shopify/sarama"
	"log")

type AppService interface {
	SendCoords([]byte) error
}

type appService struct {
	container container.Container
}

func NewAppService(container container.Container) AppService {
	return &appService{
		container: container,
	}
}


func (t *appService) SendCoords(coords []byte) error {
	topic := "coordinates"
	log.Printf("Try send location to topic %v", topic)
	addr := []string{t.container.GetConfig().KafkaHost+":"+t.container.GetConfig().KafkaPort}
	PushMessage(addr, topic, coords)
	return nil
}


func PushMessage(addrs []string, topic string, message []byte){
	producer, err := sarama.NewSyncProducer(addrs, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	msg := &sarama.ProducerMessage{Topic: topic, Value: sarama.ByteEncoder(message)}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
}

// constructor Consumer -> NewConsumer!
func StartConsume(cgroup string, topic string, zookeeperAddr string) {
    cg, err := initConsumer(cgroup, topic, zookeeperAddr)
    if err != nil {
        log.Println("Error consumer group: ", err.Error())
        os.Exit(1)
    }
    defer cg.Close()
    consume(topic, cg)
}
