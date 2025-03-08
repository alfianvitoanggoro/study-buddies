package date

import (
	"fmt"
	"time"
)

var months = [...]string{
	"Januari",
	"Februari",
	"Maret",
	"April",
	"Mei",
	"Juni",
	"Juli",
	"Agustus",
	"September",
	"Oktober",
	"November",
	"Desember",
}

var days = [...]string{
	"Minggu",
	"Senin",
	"Selasa",
	"Rabu",
	"Kamis",
	"Jumat",
	"Sabtu",
}

var layout = "2006-01-02 15:04:05"

func DateToday() *time.Time {
	now := time.Now()
	return &now
}

func DateTodayLocal() *time.Time {
	now := time.Now().UTC().Add(time.Hour * 7)
	return &now
}

func DateTodayRange() (*time.Time, *time.Time) {
	now := DateTodayLocal()
	dateStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	return &dateStart, now
}

func DateBackwardMonthRange(month int) (*time.Time, *time.Time) {
	now := DateTodayLocal()
	dateBackward := now.AddDate(0, -month, 0)
	return now, &dateBackward
}

func FormatDateIdn(t *time.Time) string {
	return fmt.Sprintf("%s, %d %s %d",
		days[t.Weekday()], t.Day(), months[t.Month()-1], t.Year())
}

func StringToTime(s string) (*time.Time, error) {
	t, err := time.Parse(layout, s)
	return &t, err
}