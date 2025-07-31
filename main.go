package main

import (
	"fmt"
	"go-events-calendar/calendar"
	"go-events-calendar/events"
)

func main() {
	e, err := events.NewEvent("Встреча", "2024-07-15 09:30")
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}

	calendar.AddEvent("event1", e)
	fmt.Println("Календарь обновлён")
	fmt.Scanln()
}
