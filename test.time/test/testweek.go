package main

import (
	"fmt"
	"time"
)

func main() {
	scheduleOfWeekday := map[time.Weekday]string{
		time.Sunday:    "",
		time.Monday:    "1:00:12",
		time.Tuesday:   "16:00:00",
		time.Wednesday: "",
		time.Thursday:  "17:00:00",
		time.Friday:    "",
		time.Saturday:  "",
	}

	t, err := getRecentEndTime(time.Now(), scheduleOfWeekday)
	fmt.Println(t, err)

}

func getRecentEndTime(now time.Time, scheduleOfWeekday map[time.Weekday]string) (time.Time, error) {
	nowWeekday := now.Weekday()
	for j := 0; j < 7; j++ {
		if nowWeekday == 7 {
			nowWeekday = nowWeekday - 7
		}
		if scheduleOfWeekday[nowWeekday] != "" {
			tOnly, err := time.Parse(time.TimeOnly, scheduleOfWeekday[nowWeekday])
			if err != nil {
				fmt.Printf("invalid ScheduleTimeSet %q %q\n", nowWeekday, scheduleOfWeekday[nowWeekday])
				return now, err
			}
			fmt.Printf("schedule time: %q\n", tOnly)
			tSchedule := time.Date(now.Year(), now.Month(), now.Day(), tOnly.Hour(), tOnly.Minute(), tOnly.Second(), 0, now.Location())
			fmt.Printf("next end schedule time: %q\n", tSchedule)
			if tSchedule.After(now) {
				return tSchedule, nil
			}
		}
		nowWeekday++
	}
	return now, fmt.Errorf("no available EndSchedule datetime")
}
