package model

type Catch struct {
	Id            int64
	HoneyId       int64
	HitTime       string
	Ip            string
	ReferralUrl   string
	XForwardedFor string
}
