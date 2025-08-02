package calendar

import (
	"fmt"
	"go-events-calendar/events"
)

var eventsMap = make(map[string]events.Event)

func AddEvent(key string, e events.Event) {
	eventsMap[key] = e
	fmt.Println("Событие добавлено:", e.Title)
}

func UpdateEvent(key string, e events.Event) {
	eventsMap[key] = e
	fmt.Println("Событие обновлено:", e.Title)
}

func ShowEvents() {
	for _, event := range eventsMap {
		fmt.Println(event.Title + " - " + event.StartAt.String())
	}
}
