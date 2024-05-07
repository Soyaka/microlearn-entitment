package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	entitment "github.com/Soyaka/entitment/api/gen"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const KafkaServer = "localhost:9092"

func main() {

	conn, err := grpc.Dial("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))

	cc := entitment.NewBookmarkServiceClient(conn)

	GetBookmarks(context.Background(), cc)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	pg := NewProducerGroup()
	defer pg.Close()

	err = pg.book.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &KafkaTopics[0], Partition: 0},
		Value:          JsData.bookByte,
	}, nil)
	handleError(err)
	time.Sleep(1 * time.Second)
	err = pg.pro.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &KafkaTopics[1], Partition: 0},
		Value:          JsData.proByte,
	}, nil)
	handleError(err)
	time.Sleep(1 * time.Second)
	err = pg.Sub.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &KafkaTopics[3], Partition: 0},
		Value:          JsData.SubByte,
	}, nil)
	handleError(err)
	time.Sleep(1 * time.Second)
	err = pg.Int.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &KafkaTopics[2], Partition: 0},
		Value:          JsData.IntByte,
	}, nil)
	handleError(err)
	time.Sleep(1 * time.Second)
}

func GetBookmarks(ctx context.Context, cc entitment.BookmarkServiceClient) {

	resp, err := cc.GetBookmark(ctx, &entitment.UserID{Id: "882e2978-80b6-4a30-938e-b0599edba0d0"})
	handleError(err)
	color.Yellow("Bookmarks: %s\n", resp.Bookmarks)
	recover()
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type dataMocker struct {
	bookmarks    *entitment.BookmarkReq
	progres      *entitment.ProgressReq
	Interest     *entitment.InterestReq
	Subscription *entitment.SubscriptionReq
}

type ProducerGroup struct {
	book *kafka.Producer
	pro  *kafka.Producer
	Sub  *kafka.Producer
	Int  *kafka.Producer
}

func NewProducerGroup() *ProducerGroup {
	return &ProducerGroup{
		book: NewProducer(),
		pro:  NewProducer(),
		Sub:  NewProducer(),
		Int:  NewProducer(),
	}
}

func NewProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": KafkaServer,
	})
	handleError(err)
	return p
}

var KafkaTopics = []string{"bookmarks", "progress", "interests", "subscriptions"}

var d = dataMocker{
	bookmarks: &entitment.BookmarkReq{
		Id:        uuid.NewString(),
		UserId:    "882e2978-80b6-4a30-938e-b0599edba0d0",
		ContentId: uuid.NewString(),
	},
	progres: &entitment.ProgressReq{
		Id:        uuid.NewString(),
		UserId:    uuid.NewString(),
		ContentId: uuid.NewString(),
		Progress:  40,
	},
	Interest: &entitment.InterestReq{
		Id:     uuid.NewString(),
		UserId: uuid.NewString(),
		Tech:   "tech",
		Cert:   "cert",
		Level:  "level",
	},
	Subscription: &entitment.SubscriptionReq{
		Id:        uuid.NewString(),
		UserId:    uuid.NewString(),
		ContentId: uuid.NewString(),
		CreatedAt: "2022-02-02T02:02:02Z",
		UpdatedAt: "2022-02-02T02:02:02Z",
	},
}

type jsonConvertor struct {
	bookByte []byte
	proByte  []byte
	SubByte  []byte
	IntByte  []byte
}

func NewJsonConvertor(intr []interface{}) *jsonConvertor {
	var jc jsonConvertor
	for _, v := range intr {
		value, err := json.Marshal(v)
		handleError(err)
		switch v.(type) {
		case *entitment.BookmarkReq:
			jc.bookByte = value
		case *entitment.ProgressReq:
			jc.proByte = value
		case *entitment.InterestReq:
			jc.IntByte = value
		case *entitment.SubscriptionReq:
			jc.SubByte = value
		}
	}
	return &jc
}

var JsData = NewJsonConvertor([]interface{}{d.bookmarks, d.progres, d.Interest, d.Subscription})

func (pg *ProducerGroup) Close() {
	pg.book.Close()
	pg.pro.Close()
	pg.Sub.Close()
	pg.Int.Close()
}
