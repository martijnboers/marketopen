package main

import (
	"fmt"
	"time"
)

type Market struct {
	Name             string
	TimeZoneLocation string
	TradingHours     []string
}

func nowToHours(timezone time.Time, timeToBeParsed string) time.Time {
	parsedTime, _ := time.Parse("15:04", timeToBeParsed)
	return time.Date(timezone.Year(), timezone.Month(), timezone.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, timezone.Location())
}

func (m Market) isOpen() bool {
	location, err := time.LoadLocation(m.TimeZoneLocation)
	if err != nil {
		fmt.Println(err)
	}
	var nowIn = time.Now().In(location)

	// Opening hours always come in pairs
	for i := 0; i < len(m.TradingHours); i += 2 {
		opening := nowToHours(nowIn, m.TradingHours[i])
		closing := nowToHours(nowIn, m.TradingHours[i+1])

		if nowIn.After(opening) && nowIn.Before(closing) {
			return true
		}
	}

	return false
}

func main() {
	markets := []Market{
		{"New York Stock Exchange (NYSE)", "America/New_York", []string{"09:30", "16:00"}},
		{"Tokyo Stock Exchange", "Asia/Tokyo", []string{"09:00", "11:30", "12:30", "15:00"}},
		{"Stock Exchange of Hong Kong (SEHK)", "Asia/Hong_Kong", []string{"09:30", "12:00", "13:00", "16:00"}},
		{"National Stock Exchange of India (NSE)", "Asia/Kolkata", []string{"09:15", "15:30"}},
		{"Shanghai Stock Exchange (SSE)", "Asia/Shanghai", []string{"09:30", "11:30", "13:00", "14:57"}},
		{"Shenzhen Stock Exchange (SZSE)", "Asia/Shanghai", []string{"09:30", "11:30", "13:00", "14:57"}},
		{"BSE Limited", "Asia/Kolkata", []string{"09:15", "15:30"}},
		{"Toronto Stock Exchange (TSX)", "America/Toronto", []string{"09:30", "16:00"}},
		{"Nasdaq Stock Market", "America/New_York", []string{"09:30", "16:00"}},
		{"London Stock Exchange", "Europe/London", []string{"08:00", "16:30"}},
		{"Frankfurt Stock Exchange", "Europe/Berlin", []string{"08:00", "22:00"}},
		{"SIX Swiss Exchange", "Europe/Zurich", []string{"09:00", "17:20"}},
		{"Euronext Amsterdam", "Europe/Amsterdam", []string{"09:00", "17:30"}},
		{"Stockholm Stock Exchange", "Europe/Stockholm", []string{"09:00", "17:25"}},
		{"B3 S.A.", "America/Sao_Paulo", []string{"10:00", "17:55"}},
		{"Johannesburg Stock Exchange (JSE)", "Africa/Johannesburg", []string{"09:00", "17:00"}},
		{"Australian Securities Exchange (ASX)", "Australia/Sydney", []string{"10:00", "16:00"}},
	}

	for _, market := range markets {
		var isOpen = "CLOSED"
		if market.isOpen() {
			isOpen = "OPEN"
		}
		fmt.Printf("%s is %s \n", market.Name, isOpen)
	}
}
