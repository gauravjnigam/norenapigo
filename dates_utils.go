package norenapigo

import (
	"fmt"
	"time"

	"github.com/golang/glog"
)

func isHoliday(today time.Time) bool {
	holidays := []string{
		"1/26/2023",
		"3/7/2023",
		"3/30/2023",
		"4/4/2023",
		"4/7/2023",
		"4/14/2023",
		"4/21/2023",
		"5/1/2023",
		"6/28/2023",
		"8/15/2023",
		"9/19/2023",
		"10/2/2023",
		"10/24/2023",
		"11/14/2023",
		"11/27/2023",
		"12/25/2023",
	}
	datesList := make([]time.Time, 0)
	for _, holiday := range holidays {
		hDate, err := time.Parse("1/2/2006", holiday)
		if err != nil {
			fmt.Println("Invalid holiday format:", holiday)
			continue
		}
		datesList = append(datesList, hDate)
	}
	// todayDate, err := time.Parse("1/2/2006", today)
	// if err != nil {
	// 	fmt.Println("Invalid today format:", today)
	// 	return false
	// }
	for _, date := range datesList {
		if today.Equal(date) {
			return true
		}
	}
	return false
}

func weeklyExpiry(dateInput time.Time) time.Time {
	weekday := int(dateInput.Weekday())
	expiry1 := dateInput.AddDate(0, 0, (10-weekday)%7)
	for isHoliday(expiry1) {
		expiry1 = expiry1.AddDate(0, 0, -1)
	}
	return expiry1
}

func monthlyExpiry(dateInput time.Time) time.Time {
	year, month, _ := dateInput.Date()
	lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()

	expiry := time.Date(year, month, lastDay, 0, 0, 0, 0, time.UTC)
	for expiry.Weekday() != time.Thursday {
		expiry = expiry.AddDate(0, 0, -1)
	}
	for isHoliday(expiry) {
		expiry = expiry.AddDate(0, 0, -1)
	}

	return expiry
}

func GetTime1(timeString string) string {
	layout := "02-01-2006 15:04:05"
	time_location, _ := time.LoadLocation("Asia/Kolkata")
	t, err := time.ParseInLocation(layout, timeString, time_location)
	// t, err := time.Parse(layout, timeString)
	if err != nil {
		glog.Fatal(err)
	}
	return fmt.Sprintf("%d", t.Unix())
}
