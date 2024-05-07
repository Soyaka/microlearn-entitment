package cache

import (
	"encoding/json"
	"time"

	entitment "github.com/Soyaka/entitment/api/gen"
	"github.com/fatih/color"
	"github.com/go-redis/redis"
)

var CACHE_TIMEOUT = 10 * time.Minute

type Cache struct {
	redis *redis.Client
}

func NewCache() *Cache {

	client := redis.NewClient(&redis.Options{

		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	color.Green("Connected to Redis\n", client)
	return &Cache{redis: client}

}

func (c *Cache) SetInterest(userId string, interest *entitment.InterestRes) error {
	tagJson, err := json.Marshal(interest)
	if err != nil {
		return err
	}
	return c.redis.Set("interest_"+userId, tagJson, CACHE_TIMEOUT).Err()
}

func (c *Cache) GetInterest(userId string) (*entitment.InterestRes, error) {
	subJson, err := c.redis.Get("interest_" + userId).Result()
	if err != nil {
		return nil, err
	}
	var resp entitment.InterestRes
	err = json.Unmarshal([]byte(subJson), &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Similarly update the Set and Get methods for other types of data (progress, subscription, bookmark) using unique key prefixes.

// define type for progres cache
type Prog struct {
	ContentId string `json:"content_id"`
	Progress  string `json:"progress"`
}

func (c *Cache) SetProgress(userId string, progress []*entitment.ProgressRes) error {

	tagJson, err := json.Marshal(progress)

	if err != nil {
		return err
	}

	return c.redis.Set("progress_"+userId, tagJson, CACHE_TIMEOUT).Err()
}

func (c *Cache) GetProgress(userId string) (*entitment.ProgressListRes, error) {

	subJson, err := c.redis.Get("progress_" + userId).Result()

	if err != nil {
		return nil, err
	}

	var res []*entitment.ProgressRes

	err = json.Unmarshal([]byte(subJson), &res)

	if err != nil {
		return nil, err
	}

	return &entitment.ProgressListRes{Progress: res}, nil
}

func (c *Cache) SetSubscription(userId string, subscription []string) error {

	tagJson, err := json.Marshal(subscription)

	if err != nil {
		return err
	}

	return c.redis.Set("subscription_"+userId, tagJson, CACHE_TIMEOUT).Err()
}

func (c *Cache) GetSubscription(userId string) (*entitment.SubscriptionRes, error) {

	subJson, err := c.redis.Get("subscription_" + userId).Result()

	if err != nil {
		return nil, err
	}

	var res []string
	err = json.Unmarshal([]byte(subJson), &res)

	if err != nil {
		return nil, err
	}

	return &entitment.SubscriptionRes{Subs: res}, nil
}

func (c *Cache) SetBookmark(userId string, bookmark []string) error {

	tagJson, err := json.Marshal(bookmark)

	if err != nil {
		return err
	}

	return c.redis.Set("bookmark_"+userId, tagJson, CACHE_TIMEOUT).Err()
}

func (c *Cache) GetBookmark(userId string) (*entitment.BookmarkRes, error) {

	subJson, err := c.redis.Get("bookmark_" + userId).Result()

	if err != nil {
		return nil, err
	}

	var res []string
	err = json.Unmarshal([]byte(subJson), &res)

	if err != nil {
		return nil, err
	}

	return &entitment.BookmarkRes{Bookmarks: res}, nil
}

func (c *Cache) DeleteInterestCache(key string) error {
	return c.redis.Del("interest_" + key).Err()
}

func (c *Cache) DeleteProgressCache(key string) error {
	return c.redis.Del("progress_" + key).Err()
}

func (c *Cache) DeleteSubscriptionCache(key string) error {
	return c.redis.Del("subscription_" + key).Err()
}

func (c *Cache) DeleteBookmarkCache(key string) error {
	return c.redis.Del("bookmark_" + key).Err()
}
