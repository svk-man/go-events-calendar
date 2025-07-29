package main

import (
	"github.com/svk-man/go-events-calendar/calendar"
	"github.com/svk-man/go-events-calendar/events"
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
