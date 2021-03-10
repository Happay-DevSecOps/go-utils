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
	Year  string
	Month string
	Day   string
	Yest  string
	Tms   int64
	Yms   int64
}

// //GetDates ...
// type GetDates interface {
// 	Ymdyesterday()
// }

//Ymdyesterday ...
func (getDate GetDate) Ymdyesterday() (getDateReturn GetDate) {
	y, m, d := currentTime.Date()
	s := d - 1

	ymsd := time.Date(y, m, d, 0, 0, 0, 0, time.UTC).UnixNano() / 1000000
	tmsd := time.Date(y, m, s, 0, 0, 0, 0, time.UTC).UnixNano() / 1000000

	getDateReturn.Year, getDateReturn.Month, getDateReturn.Day, getDateReturn.Yest, getDateReturn.Tms, getDateReturn.Yms = strconv.Itoa(y), strconv.Itoa(int(m)), strconv.Itoa(d), strconv.Itoa(s), ymsd, tmsd

	return getDateReturn
}
