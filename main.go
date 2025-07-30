package main

import (
	"go-events-calendar/calendar"
	"go-events-calendar/events"
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
	fmt.Scanln()
}
