package main

import (
	"app/calendar"
	"app/events"
	"fmt"
	"time"
)

func main() {
	e := events.Event{
		Title:   "Встреча",
		StartAt: time.Now(),
	}
	calendar.AddEvent("event1", e)
	fmt.Println("Календарь обновлён")
}
