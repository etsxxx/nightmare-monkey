package nightmare

import (
	"math/rand"
	"time"
)

// ErrorResponse error response
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var ran = rand.New(rand.NewSource(time.Now().Unix()))

type SimpleTime struct {
	Hour   int
	Minute int
}

var weekdayStrings = map[int]string{
	0: "Sun",
	1: "Mon",
	2: "Tue",
	3: "Wed",
	4: "Thu",
	5: "Fri",
	6: "Sat",
	7: "Sun",
}

func (monkey *NightmareMonkey) isSleeping() bool {
	now := time.Now().Local()

	weekday := int(now.Weekday())
	if weekday < monkey.playingDayStart || monkey.playingDayEnd < weekday {
		return true
	}

	current := now.Hour()*60 + now.Minute()
	playingTimeStart := monkey.playingTimeStart.Hour*60 + monkey.playingTimeStart.Minute
	playingTimeEnd := monkey.playingTimeEnd.Hour*60 + monkey.playingTimeEnd.Minute

	if current < playingTimeStart || playingTimeEnd < current {
		return true
	}
	return false
}
