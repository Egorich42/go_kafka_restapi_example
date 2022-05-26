package service

import ("github.com/Egorich42/testserver/container"
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


// MakePayment publishes to kafka
func (t *appService) SendCoords(coords []byte) error {
	log.Print("i try to send location")
	//message := 
	addr := []string{t.container.GetConfig().Host+":"+t.container.GetConfig().KafkaPort}
	topic := "coordinates"
	PushMessage(addr, topic, coords)
	// PUSH TO KAFKA
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

func StartConsume(cgroup string, topic string, zookeeperAddr string) {
    cg, err := initConsumer(cgroup, topic, zookeeperAddr)
    if err != nil {
        log.Println("Error consumer goup: ", err.Error())
        os.Exit(1)
    }
    defer cg.Close()

    consume(topic, cg)
}
