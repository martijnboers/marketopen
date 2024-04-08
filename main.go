package main

import (
	"fmt"
	"time"
)

type Market struct {
	Name         string
	Location     string
	TimeZone     string
	TradingHours []time.Time
}

func main() {
	markets := []Market{
		{"New York Stock Exchange (NYSE)", "New York, United States", "America/New_York", parseTimes("09:30", "16:00")},
		{"Tokyo Stock Exchange", "Tokyo, Japan", "Asia/Tokyo", parseTimes("09:00", "11:30", "12:30", "15:00")},
		{"Stock Exchange of Hong Kong (SEHK)", "Hong Kong", "Asia/Hong_Kong", parseTimes("09:30", "12:00", "13:00", "16:00")},
		{"National Stock Exchange of India (NSE)", "Mumbai, India", "Asia/Kolkata", parseTimes("09:15", "15:30")},
		{"Shanghai Stock Exchange (SSE)", "Shanghai, China", "Asia/Shanghai", parseTimes("09:30", "11:30", "13:00", "14:57")},
		{"Shenzhen Stock Exchange (SZSE)", "Shenzhen, China", "Asia/Shanghai", parseTimes("09:30", "11:30", "13:00", "14:57")},
		{"BSE Limited", "Mumbai, India", "Asia/Kolkata", parseTimes("09:15", "15:30")},
		{"Toronto Stock Exchange (TSX)", "Toronto, Canada", "America/Toronto", parseTimes("09:30", "16:00")},
		{"Nasdaq Stock Market", "New York, United States", "America/New_York", parseTimes("09:30", "16:00")},
		{"London Stock Exchange", "London, United Kingdom", "Europe/London", parseTimes("08:00", "16:30")},
		{"Frankfurt Stock Exchange", "Frankfurt, Germany", "Europe/Berlin", parseTimes("08:00", "22:00")},
		{"SIX Swiss Exchange", "Zurich, Switzerland", "Europe/Zurich", parseTimes("09:00", "17:20")},
		{"Euronext Amsterdam", "Amsterdam, Netherlands", "Europe/Amsterdam", parseTimes("09:00", "17:30")},
		{"Stockholm Stock Exchange", "Stockholm, Sweden", "Europe/Stockholm", parseTimes("09:00", "17:25")},
		{"B3 S.A.", "São Paulo, Brazil", "America/Sao_Paulo", parseTimes("10:00", "17:55")},
		{"Johannesburg Stock Exchange (JSE)", "Johannesburg, South Africa", "Africa/Johannesburg", parseTimes("09:00", "17:00")},
		{"Australian Securities Exchange (ASX)", "Sydney, Australia", "Australia/Sydney", parseTimes("10:00", "16:00")},
	}

	for _, market := range markets {
		fmt.Printf("%s := NewMarket(\"%s\", \"%s\", %v)\n", market.Name, market.Location, market.TimeZone, market.TradingHours)
	}
}

func parseTimes(times ...string) []time.Time {
	var parsedTimes []time.Time
	for _, t := range times {
		parsedTime, _ := time.Parse("15:04", t)
		parsedTimes = append(parsedTimes, parsedTime)
	}
	return parsedTimes
}

