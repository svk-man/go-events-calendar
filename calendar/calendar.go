package calendar

import (
	"github.com/svk-man/go-events-calendar/events"
	"fmt"
)

var eventsMap = make(map[string]events.Event)

func AddEvent(key string, e events.Event) {
	eventsMap[key] = e
	fmt.Println("Событие добавлено:", e.Title)
}
