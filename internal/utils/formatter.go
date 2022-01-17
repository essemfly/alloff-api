package utils

import (
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
)

func PriceFormatter(amount int) string {
	return humanize.Comma(int64(amount)) + "원"
}

func TimeFormatterForKorea(date time.Time) string {
	// MM월 DD일 A요일 HH:MM AM/PM
	loc, _ := time.LoadLocation("Asia/Seoul")
	korDate := date.In(loc)

	month := int(korDate.Month())
	day := korDate.Day()
	weekday := korDate.Weekday()
	korWeekday := ""
	switch weekday {
	case time.Monday:
		korWeekday = "월요일"
	case time.Tuesday:
		korWeekday = "화요일"
	case time.Wednesday:
		korWeekday = "수요일"
	case time.Thursday:
		korWeekday = "목요일"
	case time.Friday:
		korWeekday = "금요일"
	case time.Saturday:
		korWeekday = "토요일"
	case time.Sunday:
		korWeekday = "일요일"
	}
	hour := korDate.Hour()
	minute := korDate.Minute()
	minuteStr := strconv.Itoa(minute)
	if len(minuteStr) == 1 {
		minuteStr = "0" + minuteStr
	}
	ampm := "AM"
	if hour >= 13 {
		ampm = "PM"
		hour -= 12
	} else if hour == 12 {
		ampm = "PM"
	}

	return strconv.Itoa(month) + "월 " + strconv.Itoa(day) + "일 " + korWeekday + " " + strconv.Itoa(hour) + ":" + minuteStr + " " + ampm
}
