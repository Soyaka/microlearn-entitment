package handlers

import (
	"context"
	"time"

	entitment "github.com/Soyaka/entitment/api/gen"
	"github.com/Soyaka/entitment/internal/cache"
	"github.com/Soyaka/entitment/internal/database"
	"github.com/Soyaka/entitment/internal/kafka"
)

type GlobalHandler struct {
	Data  *database.Service
	Cache *cache.Cache
	Kafka *kafka.Consumer

	entitment.UnimplementedBookmarkServiceServer
	entitment.UnimplementedInterestServiceServer
	entitment.UnimplementedProgressServiceServer
	entitment.UnimplementedSubscriptionServiceServer
}

var TimeOut = 100 * time.Millisecond

func NewGlobalHandler() *GlobalHandler {
	return &GlobalHandler{
		Data:  database.NewService(),
		Cache: cache.NewCache(),
		Kafka: kafka.NewConsumer(),
	}
}

func (g *GlobalHandler) GetBookmark(ctx context.Context, in *entitment.UserID) (*entitment.BookmarkRes, error) {
	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	var bookmarks *entitment.BookmarkRes

	bookmarks, err := g.Cache.GetBookmark(in.Id)
	if err != nil || bookmarks == nil {
		bookmarks, err = g.Data.GetBookmarks(ctx, in.Id)
		if err != nil {
			return nil, err
		}
		g.Cache.SetBookmark(in.Id, bookmarks.Bookmarks)
	}
	return bookmarks, nil
}

func (g *GlobalHandler) GetInterests(ctx context.Context, in *entitment.UserID) (*entitment.InterestRes, error) {
	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	var interests *entitment.InterestRes

	interests, err := g.Cache.GetInterest(in.Id)
	if err != nil || interests == nil {
		interests, err = g.Data.GetInterests(ctx, in.Id)
		if err != nil {
			return nil, err
		}
		g.Cache.SetInterest(in.Id, interests)
	}

	return interests, nil

}

func (g *GlobalHandler) GetProgress(ctx context.Context, in *entitment.UserID) (*entitment.ProgressListRes, error) {

	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	var res *entitment.ProgressListRes
	res, err := g.Cache.GetProgress(in.Id)
	if err != nil || res == nil {

		res, err = g.Data.GetProgress(ctx, in.Id)
		if err != nil {
			return nil, err
		}
		g.Cache.SetProgress(in.Id, res.Progress)
	}

	return res, nil

}

func (g *GlobalHandler) GetSubscription(ctx context.Context, in *entitment.UserID) (*entitment.SubscriptionRes, error) {

	ctx, cancel := context.WithTimeout(ctx, TimeOut)

	defer cancel()
	var res *entitment.SubscriptionRes
	res, err := g.Cache.GetSubscription(in.Id)

	if err != nil || res == nil {
		res, err := g.Data.GetSubscriptions(ctx, in.Id)
		if err != nil {
			return nil, err
		}
		g.Cache.SetSubscription(in.Id, res.Subs)
	}
	return res, nil

}

var (
	KafkaServer  = "localhost:9092"
	KafkaTopics  = []string{"bookmarks", "interests", "progress", "subscriptions"}
	KafkaGroupId = "rdkafka-f9a5f53b-cc4e-45b403cee"
)

//TODO: Add a config object inside the golbal handler and put all the needded variables here
