package nightmare

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const appName = "Nightmare Monkey"

type NightmareMonkey struct {
	Host             string
	Name             string
	ListenPort       int
	Logger           *log.Logger
	Interval         uint64
	Dryrun           bool
	playingDayStart  int
	playingDayEnd    int
	playingTimeStart SimpleTime
	playingTimeEnd   SimpleTime
}

func New(port int, interval uint64, playingDay string, playingTime string) (monkey *NightmareMonkey, err error) {
	days := strings.SplitN(playingDay, "-", 2)
	startDay, err := strconv.Atoi(days[0])
	if err != nil {
		return nil, fmt.Errorf("day format error")
	}
	endDay, err := strconv.Atoi(days[1])
	if err != nil {
		return nil, fmt.Errorf("day format error")
	}

	times := strings.SplitN(playingTime, "-", 2)
	startTimes := strings.SplitN(times[0], ":", 2)
	startHour, err := strconv.Atoi(startTimes[0])
	if err != nil {
		return nil, fmt.Errorf("time format error")
	}
	startMinute, err := strconv.Atoi(startTimes[1])
	if err != nil {
		return nil, fmt.Errorf("time format error")
	}

	endTimes := strings.SplitN(times[1], ":", 2)
	endHour, err := strconv.Atoi(endTimes[0])
	if err != nil {
		return nil, fmt.Errorf("time format error")
	}
	endMinute, err := strconv.Atoi(endTimes[1])
	if err != nil {
		return nil, fmt.Errorf("time format error")
	}

	host, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	monkey = &NightmareMonkey{
		Name:             appName,
		Host:             host,
		Interval:         interval,
		ListenPort:       port,
		Logger:           log.New(os.Stdout, "", log.LstdFlags),
		playingDayStart:  startDay,
		playingDayEnd:    endDay,
		playingTimeStart: SimpleTime{Hour: startHour, Minute: startMinute},
		playingTimeEnd:   SimpleTime{Hour: endHour, Minute: endMinute},
	}

	monkey.Logger.Printf("Nightmare Monkey is playing from %02d:%02d to %02d:%02d, %s thru %s.",
		monkey.playingTimeStart.Hour, monkey.playingTimeStart.Minute,
		monkey.playingTimeEnd.Hour, monkey.playingTimeEnd.Minute,
		weekdayStrings[monkey.playingDayStart],
		weekdayStrings[monkey.playingDayEnd],
	)

	return monkey, nil
}
