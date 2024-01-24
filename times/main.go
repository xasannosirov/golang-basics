package main

import (
	"fmt"
	"time"
)

func TimeDate() {
	current := time.Date(2024, time.May, 14, 12, 45, 12, 0, time.Local)
	fmt.Println(current.Date())
	fmt.Println(current.Year())
	fmt.Println(current.Month())
	fmt.Println(current.Day())
	fmt.Println(current.Clock())
	fmt.Println(current.Minute())
	fmt.Println(current.Second())
	fmt.Println(current.Unix())
	fmt.Println(current.Weekday())
	fmt.Println(current.YearDay())
}

func TimeFormat() {
	current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	fmt.Println(current.Format("02-01-2006 15:04:05"))
}

func TimeCheck() {
	firstTime := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	secondTime := time.Date(2020, time.May, 15, 16, 45, 12, 0, time.Local)

	fmt.Println(firstTime.After(secondTime))
	fmt.Println(firstTime.Before(secondTime))
	fmt.Println(firstTime.Equal(secondTime))
}

func TimeMethods() {
	now := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

	future := now.Add(time.Hour * 12)
	past := now.AddDate(-1, -2, -3)
	fmt.Println(future.Sub(past))
}

func TimeDuration() {
	now := time.Now()
	past := now.AddDate(0, 0, -1)
	future := now.AddDate(0, 0, 1)

	fmt.Println(time.Since(past).Round(time.Second))
	fmt.Println(time.Until(future).Round(time.Second))
	dur, err := time.ParseDuration("1h12m3s")
	if err != nil {
		panic(err)
	}
	fmt.Println(dur.Round(time.Hour).Hours()) // 1
}

func main() {
	TimeDate()
	TimeFormat()
	TimeCheck()
	TimeMethods()
}
