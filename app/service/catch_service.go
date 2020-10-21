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
		HoneyId:       record[0].Id,
		HitTime:       hitTime,
		Ip:            ip,
		ReferralUrl:   referralUrl,
		XForwardedFor: xForwardedFor,
	}

	ipLookup, err := GetIpData(ip)
	if err != nil {
		log.Println(err)
	} else {
		c.City = ipLookup.City
		c.CountryName = ipLookup.CountryName
		c.CountryCode = ipLookup.CountryCode
		c.ContinentName = ipLookup.ContinentName
		c.Latitude = ipLookup.Latitude
		c.Longitude = ipLookup.Longitude
		c.AsnId = ipLookup.ASN.ASN
		c.AsnName = ipLookup.ASN.Name
		c.AsnDomain = ipLookup.ASN.Domain
		c.AsnRoute = ipLookup.ASN.Route
		c.AsnType = ipLookup.ASN.Type
		c.IsTor = ipLookup.Threat.IsTOR
		c.IsProxy = ipLookup.Threat.IsProxy
		c.IsAnonymous = ipLookup.Threat.IsAnonymous
		c.IsKnownAttacker = ipLookup.Threat.IsKnownAttacker
		c.IsKnownAbUser = ipLookup.Threat.IsKnownAbuser
		c.IsThreat = ipLookup.Threat.IsThreat
		c.IsBogon = ipLookup.Threat.IsBogon
	}

	return repository.CreateCatch(c)
}
