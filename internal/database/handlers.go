package database

import (
	"context"
	"time"

	entitment "github.com/Soyaka/entitment/api/gen"
	"github.com/fatih/color"
)

var TimeOut = 100 * time.Millisecond

func (s *Service) GetSubscriptions(ctx context.Context, id string) (*entitment.SubscriptionRes, error) {
	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	var subscriptions []entitment.SubscriptionReq
	res := s.subscriptions.Find(&subscriptions, "user_id = ?", id).WithContext(ctx)

	if res.Error != nil {
		return nil, res.Error
	}

	var SubRes []string	

	for _, sub := range subscriptions {
		SubRes = append(SubRes, sub.ContentId)
	}

	return &entitment.SubscriptionRes{Subs: SubRes}, nil
}

func (s *Service) GetBookmarks(ctx context.Context, id string) (*entitment.BookmarkRes, error) {
	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	var Bookmarks *[]entitment.BookmarkReq
	res := s.bookmarks.Find(&Bookmarks, "user_id = ?", id).WithContext(ctx)

	if res.Error != nil {
		return nil, res.Error
	}

	var BookRes []string

	for _, bookmark := range *Bookmarks {
		BookRes = append(BookRes, bookmark.ContentId)
	}

	return &entitment.BookmarkRes{Bookmarks: BookRes}, nil
}

func (s *Service) GetProgress(ctx context.Context, id string) (*entitment.ProgressListRes, error) {
	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	var Progress *[]entitment.ProgressReq
	res := s.progress.Find(&Progress, "user_id = ?", id).WithContext(ctx)

	if res.Error != nil {
		return nil, res.Error
	}
	var Progresses []*entitment.ProgressRes

	for _, progress := range *Progress {
		pro := &entitment.ProgressRes{ContentId: progress.ContentId, Progress: progress.Progress}

		Progresses = append(Progresses, pro)
	}

	return &entitment.ProgressListRes{Progress: Progresses}, nil
}

func (s *Service) GetInterests(ctx context.Context, id string) (*entitment.InterestRes, error) {
	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()
	var Interests *entitment.InterestReq
	res := s.interests.Find(&Interests, "user_id = ?", id).WithContext(ctx)
	if res.Error != nil {
		return nil, res.Error
	}

	return &entitment.InterestRes{Tech: Interests.Tech, Cert: Interests.Cert, Level: Interests.Level}, nil
}

// add new things to database that will  handled by kafka
func (s *Service) CreateBookmark(ctx context.Context, bookmark *entitment.BookmarkReq) error {

	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	return s.bookmarks.Create(&bookmark).WithContext(ctx).Error
}

func (s *Service) CreateSubscription(ctx context.Context, subscription *entitment.SubscriptionReq) error {
	color.Red("Subscription: %s\n", subscription)
	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	return s.subscriptions.Create(&subscription).WithContext(ctx).Error

}

func (s *Service) CreateProgress(ctx context.Context, progress *entitment.ProgressReq) error {
	color.Red("Progress: %s\n", progress)
	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	return s.progress.Create(&progress).WithContext(ctx).Error

}

func (s *Service) CreateInterest(ctx context.Context, interest *entitment.InterestReq) error {

	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	return s.interests.Create(&interest).WithContext(ctx).Error
}
