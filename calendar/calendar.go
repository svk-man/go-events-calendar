package calendar

import (
	"fmt"
	"go-events-calendar/events"
)

var calendarEvents = make(map[string]events.Event)

func AddEvent(title string, date string) (events.Event, error) {
	e, err := events.NewEvent(title, date)
	if err != nil {
		return events.Event{}, err
	}

	calendarEvents[e.ID] = e
	fmt.Println("Событие добавлено:", e.ID)

	return e, nil
}

func EditEvent(id string, title string, date string) error {
	e, err := events.NewEvent(title, date)
	if err != nil {
		return err
	}

	e.ID = id
	calendarEvents[e.ID] = e

	fmt.Println("Событие обновлено:", e.ID)

	return nil
}

func DeleteEvent(id string) {
	targetKey := ""
	for key, event := range calendarEvents {
		if event.ID == id {
			targetKey = key
		}
	}

	if targetKey != "" {
		delete(calendarEvents, targetKey)
		fmt.Println("Событие удалено:", id)
	}
}

func ShowEvents() {
	for _, event := range calendarEvents {
		fmt.Println(event.Title + " - " + event.StartAt.String())
	}
}
