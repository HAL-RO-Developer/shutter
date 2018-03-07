package service

import (
	"github.com/HAL-RO-Developer/shutter/model"
)

var Account = accountService{}

type accountService struct {
}

func (a *accountService) Get(twitterID string) *model.Account {
	var accounts []model.Account
	db.Where("twitter_id = ?", twitterID).Find(&accounts)
	if len(accounts) == 0 {
		return nil
	}
	return &accounts[0]
}
