package handlers

import (
	"context"
	"encoding/json"
	"log"

	entitment "github.com/Soyaka/microlearn-entitment/api/gen"
	"github.com/fatih/color"
)

func (g *GlobalHandler) MessageReader() {

	for {
		msg, err := g.Kafka.Consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
		log.Println("Consumed message:", string(msg.Value))
		topic := string(*msg.TopicPartition.Topic)
		color.Green("Consumed from topic: %s\n", topic)
		switch string(*msg.TopicPartition.Topic) {
		case KafkaTopics[0]: //// bookmarks
			log.Printf("Bookmarks: %s\n", msg.Value)
			err = g.BookmarkMsgResponder(context.Background(), msg.Value)
			if err != nil {
				color.Red("Error: %v\n", err)
			}
		case KafkaTopics[1]: //// interests
			log.Printf("Interests: %s\n", msg.Value)
			err = g.InterestMsgResponder(context.Background(), msg.Value)
			if err != nil {
				color.Red("Error: %v\n", err)
			}
		case KafkaTopics[2]: //// progress
			log.Printf("Progress: %s\n", msg.Value)
			err = g.ProgressMsgResponder(context.Background(), msg.Value)
			if err != nil {
				color.Red("Error: %v\n", err)
			}
		case KafkaTopics[3]: //// subscriptions
			log.Printf("Subscriptions: %s\n", msg.Value)
			err = g.SubscriptionMsgResponder(context.Background(), msg.Value)
			if err != nil {
				color.Red("Error: %v\n", err)
			}
		}
	}
}

func (g *GlobalHandler) BookmarkMsgResponder(ctx context.Context, fill []byte) error {
	ctxx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	var cn *entitment.BookmarkReq
	err := json.Unmarshal(fill, &cn)
	if err != nil {
		color.Red("Error: %v\n", err)//TODO: go back later to handle error here Mustn't be fatal
	}
	g.Cache.DeleteBookmarkCache(cn.UserId)
	return g.Data.CreateBookmark(ctxx, cn)
}

func (g *GlobalHandler) InterestMsgResponder(ctx context.Context, fill []byte) error {
	ctxx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	var cn *entitment.InterestReq
	err := json.Unmarshal(fill, &cn)
	if err != nil {
		color.Red("Error: %v\n", err) //TODO: go back later to handle error here Mustn't be fatal
	}
	g.Cache.DeleteInterestCache(cn.UserId)
	return g.Data.CreateInterest(ctxx, cn)
}

func (g *GlobalHandler) ProgressMsgResponder(ctx context.Context, fill []byte) error {
	ctxx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	var cn *entitment.ProgressReq
	err := json.Unmarshal(fill, &cn)
	if err != nil {
		color.Red("Error: %v\n", err) //TODO: go back later to handle error here Mustn't be fatal
	}
	g.Cache.DeleteProgressCache(cn.UserId)
	return g.Data.CreateProgress(ctxx, cn)
}

func (g *GlobalHandler) SubscriptionMsgResponder(ctx context.Context, fill []byte) error {
	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	color.Blue("SUBS: %s\n", fill)
	var cn *entitment.SubscriptionReq
	err := json.Unmarshal(fill, &cn)
	if err != nil {
		log.Fatal(err) //TODO: go back later to handle error here Mustn't be fatal
	}
	go g.Cache.DeleteSubscriptionCache(cn.UserId)
	return g.Data.CreateSubscription(ctx, cn)
}
