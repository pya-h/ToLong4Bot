package main

import (
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) Get() string {

	return fmt.Sprintf("%d-%d-%d\t%d:%d", d.Year(), d.Month(), d.Day(), d.Hour(), d.Minute())
}

func (d *Date) Short() string {

	return fmt.Sprintf("%d-%d-%d", d.Year(), d.Month(), d.Day())
}

// func Now() time.Time {
// if timeZone, err := time.LoadLocation("Asia/Tehran"); err == nil {
// 	return time.Now().In(timeZone)
// }
// 	return time.Now()
// }

func Today() Date {
	return Date{time.Now()}
}

func (date *Date) ToLocal() Date {
	if timeZone, err := time.LoadLocation("Asia/Tehran"); err == nil {
		return Date{date.In(timeZone)}
	}
	return *date
}

func (date *Date) ToLocal2() (res Date) {
	if locale, err := time.LoadLocation("Asia/Tehran"); err == nil {

		res = Date{time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), locale)}
	} else {
		res = Date{time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), time.Local)}

	}
	return
}
func main() {
	x := Today()
	y := x.ToLocal()
	z := x.ToLocal2()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
