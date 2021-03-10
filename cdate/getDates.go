package cdate

import (
	"strconv"
	"time"
)

var (
	currentTime time.Time = time.Now()
)

//GetDate ...
type GetDate struct {
	year  string
	month string
	day   string
	yest  string
	tms   int64
	yms   int64
}

//GetDates ...
type GetDates interface {
	ymdyesterday()
}

//Ymdyesterday ...
func (getDate GetDate) Ymdyesterday() (getDateReturn GetDate) {
	y, m, d := currentTime.Date()
	s := d - 1

	ymsd := time.Date(y, m, d, 0, 0, 0, 0, time.UTC).UnixNano() / 1000000
	tmsd := time.Date(y, m, s, 0, 0, 0, 0, time.UTC).UnixNano() / 1000000

	getDateReturn.year, getDateReturn.month, getDateReturn.day, getDateReturn.yest, getDateReturn.tms, getDateReturn.yms = strconv.Itoa(y), strconv.Itoa(int(m)), strconv.Itoa(d), strconv.Itoa(s), ymsd, tmsd

	return getDateReturn
}
