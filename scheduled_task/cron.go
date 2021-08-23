package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func taskNoParam() {
	fmt.Println("no param task: ", time.Now().UTC())
}

func stopper(s *gocron.Scheduler) {
	time.Sleep(30 * time.Second)
	fmt.Println("stopping taskNoParam...")
	s.Remove(taskNoParam)
	time.Sleep(2 * time.Minute)
	fmt.Println("stopping every tasks...")
	s.Stop()
	fmt.Println("stopper end")
}
func Cron_solution() {
	fmt.Println("Cron solution start: ")
	Task(time.Now().UTC())
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Second().Do(taskNoParam)
	alarm := time.Now().Add(time.Minute).UTC()
	str_alarm := fmt.Sprintf("%02d:%02d:00", alarm.Hour(), alarm.Minute())
	fmt.Println("scheduling main Task at: ", str_alarm)
	s.Every(1).Day().At(str_alarm).Do(Task, time.Now().UTC())
	s.StartAsync()
	stopper(s)
}
