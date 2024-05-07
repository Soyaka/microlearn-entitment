package database

import (
	"log"

	entitment "github.com/Soyaka/entitment/api/gen"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Service struct {
	bookmarks     *gorm.DB
	interests     *gorm.DB
	progress      *gorm.DB
	subscriptions *gorm.DB
}

func NewService() *Service {

	bookdb, err := gorm.Open(sqlite.Open("data/bookmarks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	interestdb, err := gorm.Open(sqlite.Open("data/interests.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	progressdb, err := gorm.Open(sqlite.Open("data/progress.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	subscriptiondb, err := gorm.Open(sqlite.Open("data/subscriptions.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err) //TODO: handle error
	}

	bookdb.AutoMigrate(&entitment.BookmarkReq{})
	progressdb.AutoMigrate(&entitment.ProgressReq{})
	interestdb.AutoMigrate(&entitment.InterestReq{})
	subscriptiondb.AutoMigrate(&entitment.SubscriptionReq{})

	return &Service{bookmarks: bookdb, interests: interestdb, progress: progressdb, subscriptions: subscriptiondb}
}
