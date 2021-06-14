package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	connType        = "tcp"
	listenIP        = "0.0.0.0"
	listenPORT      = "3380"
	tempStorage     []resConsume            // array for save consume result
	channelArrChan  = make(chan resConsume) // channel for
	producerArrChan = make(chan resConsume) // channel for receive data from `Consumer (Kafka)` and send data to channelChan
)

type resConsume struct {
	Head    string `json:"stan"`
	Content string `json:"msgin"`
}

func main() {
	go kafkaConsumer()
	l, err := net.Listen(connType, listenIP+":"+listenPORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + listenIP + ":" + listenPORT)

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("new connection detected!")
		go testReceive(conn)
		go kafkaProducer()
	}
}

func testReceive(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)

		message := make([]byte, 1024)
		_, readErr := reader.Read(message)
		message = bytes.Trim(message, "\x00")
		if readErr != nil {
			fmt.Println("failed:", readErr)
			return
		}

		head := strconv.Itoa(makeTimestamp())

		data := resConsume{
			Head:    head,
			Content: string(message),
		}

		channelArrChan <- data
		go testSend(conn, head)
	}
}

func makeTimestamp() int {
	return int(time.Now().UnixNano()) / int(time.Millisecond)
}

func testSend(conn net.Conn, head string) {
	msgConsume := getValueFromArray(head)
	var response string
	if msgConsume.Head == head {
		response = msgConsume.Content
	} else {
		response = "fail to get response"
	}
	// send to server
	fmt.Fprintf(conn, response+"\n")
	// wait for reply
	message, _ := bufio.NewReader(conn).ReadString('\n')

	fmt.Println(message)
}

func produceMsgToKafka(broker string, topic string, message <-chan resConsume) {
	//put channel data in variable
	data := <-message
	fmt.Println("Producer started!")

	// Setting up Consumer (Kafka) config
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	//set msg header
	header := map[string]string{
		"key":   "uniqueKey",
		"value": data.Head,
	}

	//produce msg
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(data.Content),
		Headers:        []kafka.Header{{Key: header["key"], Value: []byte(header["value"])}},
	}, nil)

	p.Flush(3 * 1000)

	log.Println("Producer closing!")

	// Done with worker
}

func kafkaProducer() {
	for {
		select {
		case newRequest := <-channelArrChan:
			log.Println("New request from `Channel` is ready to produce to Kafka")

			// Run Producer (Kafka)
			go produceMsgToKafka("localhost:9092", "mpc.json.bifast.request", producerArrChan)

			// Send new request to producerChan, then produce the new request to Kafka
			producerArrChan <- newRequest

		// keep looping if there is none new request
		default:
			continue
		}
	}
}

func kafkaConsumer() {

	// Setting up Consumer (Kafka) config
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "test",
		"auto.offset.reset": "latest",
	})

	if err != nil {
		log.Println(err.Error())
	}

	// Subscribe to topics
	c.SubscribeTopics([]string{"mpc.json.bifast.response"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			log.Println("New Request from Kafka")
			log.Printf("Message consumed on %s: %s\n", msg.TopicPartition, string(msg.Value))
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

			// Send any consumed event to consumerChan
			// go getResponseUpdated(consumerChan)

			msgResult := resConsume{
				Head:    string(msg.Headers[0].Value),
				Content: string(msg.Value),
			}
			tempStorage = append(tempStorage, msgResult)
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func getValueFromArray(head string) resConsume {
	search := true
	var res resConsume
	limit := 0
	for search {
		index := 0
		fmt.Println(len(tempStorage))
		for _, ele := range tempStorage {
			if ele.Head == head {
				search = false
				res = ele
				tempStorage = append(tempStorage[:index], tempStorage[index+1:]...)
			}
			index++
		}
		time.Sleep(100 * time.Millisecond)
		limit++
		if limit == 500 {
			break
		}

	}
	return res
}
