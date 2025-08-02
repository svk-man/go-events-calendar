package calendar

import (
	"errors"
	"fmt"
	"go-events-calendar/events"
)

var eventsMap = make(map[string]events.Event)

func AddEvent(key string, e events.Event) error {
	if !events.IsValidTitle(e.Title) {
		return errors.New("неверный формат заголовка")
	}

	eventsMap[key] = e
	fmt.Println("Событие добавлено:", e.Title)

	return nil
}

func UpdateEvent(key string, e events.Event) error {
	if !events.IsValidTitle(e.Title) {
		return errors.New("неверный формат заголовка")
	}

	eventsMap[key] = e
	fmt.Println("Событие обновлено:", e.Title)

	return nil
}

func RemoveEvent(key string) {
	fmt.Println("Событие удалено:", eventsMap[key].Title)
	delete(eventsMap, key)
}

func ShowEvents() {
	for _, event := range eventsMap {
		fmt.Println(event.Title + " - " + event.StartAt.String())
	}
}
