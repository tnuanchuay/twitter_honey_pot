package model

type Catch struct {
	Id              int64
	HoneyId         int64
	HitTime         string
	Ip              string
	ReferralUrl     string
	XForwardedFor   string
	City            string
	CountryName     string
	CountryCode     string
	ContinentName   string
	Latitude        float64
	Longitude       float64
	AsnId           string
	AsnName         string
	AsnDomain       string
	AsnRoute        string
	AsnType         string
	IsTor           bool
	IsProxy         bool
	IsAnonymous     bool
	IsKnownAttacker bool
	IsKnownAbUser   bool
	IsThreat        bool
	IsBogon         bool
}
