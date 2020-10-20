package service

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
	"twitter_honey_pod/app/model"
	"twitter_honey_pod/app/repository"
)

func StoreNewCatch(url, ip, referralUrl, xForwardedFor string) (*model.Catch, error) {
	url = strings.TrimLeft(url, "/")
	var hitTime = time.Now().Format("2006-01-02 15:04:05")
	h := model.Honey{
		Url: url,
	}

	record, err := repository.GetHoney(&h)
	if err != nil {
		return nil, err
	}

	if len(record) != 1 {
		log.Println("record length", len(record))
		return nil, errors.New(fmt.Sprintf("found %s more than 1 record", url))
	}

	c := model.Catch{
		HoneyId:     record[0].Id,
		HitTime:     hitTime,
		Ip:          ip,
		ReferralUrl: referralUrl,
		XForwardedFor: xForwardedFor,
	}

	return repository.CreateCatch(c)
}
