package service

import (
	"errors"
	"strings"
	"time"
	"twitter_honey_pod/app/model"
	"twitter_honey_pod/app/repository"
)

func StoreNewHoney(userId, redirectTo, url string) (*model.Honey, error){
	var createDate = time.Now().Format("2006-01-02 15:04:05")
	h := model.Honey{
		RedirectTo: redirectTo,
		UserId:     userId,
		CreateDate: createDate,
		Url: url,
	}
	result, err := repository.CreateHoney(h)
	return result, err
}

func FindHoneyByPath (path string) (*model.Honey, error) {
	path = strings.TrimLeft(path, "/")
	h := model.Honey{
		Url: path,
	}
	result, err := repository.GetHoney(&h)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("honey not found")
	}

	return &result[0], nil
}