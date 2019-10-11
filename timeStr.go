package gtool

import (
	"fmt"
	"strconv"
	"time"
)

//timestamp format year month day
func TsFormatYmd(unix int64) string {
	if unix == 0 {
		return time.Unix(time.Now().Unix(), 0).Format(FORMAT_TIME_YMD)
	} else {
		return time.Unix(unix, 0).Format(FORMAT_TIME_YMD)
	}
}

//timestamp format month day
func TsFormatMd(ts int64) string {
	t, isNow := isNowYear(ts)

	if isNow {
		return t.Format(FORMAT_TIME_MD)
	} else {
		return t.Format(FORMAT_TIME_YMD)
	}
}

//timestamp format hour minute
func TsFormatHi(ts int64) string {
	t, isNow := isNowYear(ts)

	if isNow {
		return t.Format(FORMAT_TIME_MD_HI)
	} else {
		return t.Format(FORMAT_TIME_YMD_HI)
	}
}

//accurate timestamp string
func TsFormatAccurate(ts int64) string {
	var retStr string
	nowTs := TsNow()
	offset := nowTs - ts

	if offset >= 0 { //以前
		if offset < 60 {
			retStr = strconv.FormatInt(offset, 10) + "秒前"
		} else if offset < 3600 {
			retStr = strconv.FormatInt(offset/60, 10) + "分钟前"
		} else if offset < 86400 {
			retStr = strconv.FormatInt(offset/3600, 10) + "小时前"
		} else if offset < 259200 {
			retStr = strconv.FormatInt(offset/86400, 10) + "天前"
		} else {
			retStr = TsFormatYmd(ts)
		}
	} else { //未来
		offset = ts - nowTs

		if offset < 60 {
			retStr = strconv.FormatInt(offset, 10) + "秒后"
		} else if offset < 3600 {
			retStr = strconv.FormatInt(offset/60, 10) + "分钟后"
		} else if offset < 86400 {
			retStr = strconv.FormatInt(offset/3600, 10) + "小时后"
		} else if offset < 259200 {
			retStr = strconv.FormatInt(offset/86400, 10) + "天后"
		} else {
			retStr = TsFormatYmd(ts)
		}

	}

	return retStr
}

//time string covert timestamp
func TimeStrToTs(timeStr string) (int64) {
	theTime, _ := time.ParseInLocation(FORMAT_TIME_YMD_HIS, timeStr, TimeLoc())
	return theTime.Unix()
}

//now time string
func TimeNow() string {
	return time.Now().Format(FORMAT_TIME_YMD_HIS)
}

//date string today
func DateToday() (dateStr string) {
	return DateYmd(time.Now().Unix())
}

//date string today begin time
func DateTodayBegin() string {
	return fmt.Sprintf("%s 00:00:00", DateToday())
}

//date string today end time
func DateTodayEnd() string {
	return fmt.Sprintf("%s 23:59:59", DateToday())
}

//date string yesterday
func DateYesterday() (dateStr string) {
	return DateYmdDay(-1)
}

//date string tomorrow
func DateTomorrow() (dateStr string) {
	return DateYmdDay(1)
}

//date string use days
func DateYmdDay(days int) (dateStr string) {
	ts := time.Now().AddDate(0, 0, days).Unix()
	return DateYmd(ts)
}

//date string use timestamp
func DateYmd(ts int64) (dateStr string) {
	return time.Unix(ts, 0).Format(FORMAT_TIME_YMD)
}
