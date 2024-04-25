package main

import (
	"fmt"
	"time"
)

const timeFormat = "15:04"

type Market struct {
	Name         string
	TradingHours []time.Time
	LocalTime    time.Time
}

func NewMarket(name string, timezoneLocation string, times ...string) *Market {
	var parsedTimes []time.Time
	location, err := time.LoadLocation(timezoneLocation)

	if err != nil {
		fmt.Println(err)
	}
	nowIn := time.Now().In(location)

	for _, t := range times {
		parsedTime, _ := time.Parse(timeFormat, t)
		parsedDateTime := time.Date(nowIn.Year(), nowIn.Month(), nowIn.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, nowIn.Location())
		parsedTimes = append(parsedTimes, parsedDateTime)
	}

	market := Market{name, parsedTimes, nowIn}
	return &market
}

func (m Market) isOpen() bool {
	// Opening hours always come in pairs
	for i := 0; i < len(m.TradingHours); i += 2 {
		opening := m.TradingHours[i]
		closing := m.TradingHours[i+1]

		if m.LocalTime.After(opening) && m.LocalTime.Before(closing) {
			return true
		}
	}

	return false
}

func main() {
	markets := []Market{
		*NewMarket("New York Stock Exchange (NYSE)", "America/New_York", "09:30", "16:00"),
		*NewMarket("Tokyo Stock Exchange", "Asia/Tokyo", "09:00", "11:30", "12:30", "15:00"),
		*NewMarket("Stock Exchange of Hong Kong (SEHK)", "Asia/Hong_Kong", "09:30", "12:00", "13:00", "16:00"),
		*NewMarket("National Stock Exchange of India (NSE)", "Asia/Kolkata", "09:15", "15:30"),
		*NewMarket("Shanghai Stock Exchange (SSE)", "Asia/Shanghai", "09:30", "11:30", "13:00", "14:57"),
		*NewMarket("Shenzhen Stock Exchange (SZSE)", "Asia/Shanghai", "09:30", "11:30", "13:00", "14:57"),
		*NewMarket("BSE Limited", "Asia/Kolkata", "09:15", "15:30"),
		*NewMarket("Toronto Stock Exchange (TSX)", "America/Toronto", "09:30", "16:00"),
		*NewMarket("Nasdaq Stock Market", "America/New_York", "09:30", "16:00"),
		*NewMarket("London Stock Exchange", "Europe/London", "08:00", "16:30"),
		*NewMarket("Frankfurt Stock Exchange", "Europe/Berlin", "08:00", "22:00"),
		*NewMarket("SIX Swiss Exchange", "Europe/Zurich", "09:00", "17:20"),
		*NewMarket("Euronext Amsterdam", "Europe/Amsterdam", "09:00", "17:30"),
		*NewMarket("Stockholm Stock Exchange", "Europe/Stockholm", "09:00", "17:25"),
		*NewMarket("B3 S.A.", "America/Sao_Paulo", "10:00", "17:55"),
		*NewMarket("Johannesburg Stock Exchange (JSE)", "Africa/Johannesburg", "09:00", "17:00"),
		*NewMarket("Australian Securities Exchange (ASX)", "Australia/Sydney", "10:00", "16:00"),
	}

	for _, market := range markets {
		var isOpen = "CLOSED"
		if market.isOpen() {
			isOpen = "OPEN"
		}
		fmt.Printf("%s is %s \n", market.Name, isOpen)
	}
}
