package gtool

import "time"

const (
	FORMAT_TIME_YMD_HIS = "2006-01-02 15:04:05"
	FORMAT_TIME_YMD_HI  = "2006-01-02 15:04"
	FORMAT_TIME_MD_HI   = "01-02 15:04"
	FORMAT_TIME_Y       = "2006"
	FORMAT_TIME_YMD     = "2006-01-02"
	FORMAT_TIME_MD      = "01-02"
	FORMAT_TIME_HI      = "15:04"
	FORMAT_TIME_HIS     = "15:04:05"
)

//time location
func TimeLoc() *time.Location {
	loc, _ := time.LoadLocation("Local")
	return loc
}

//timestamp now
func TsNow() int64 {
	return time.Now().Unix()
}

//timestamp is now year
func isNowYear(ts int64) (retTime time.Time, isNow bool) {
	t := time.Unix(ts, 0)
	toNow := time.Unix(time.Now().Unix(), 0)
	nowYear := toNow.Format(FORMAT_TIME_Y)
	timeYear := t.Format(FORMAT_TIME_Y)

	if nowYear == timeYear {
		return t, true
	}
	return t, false
}
