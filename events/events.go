package events

import (
	"errors"
	"time"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID      string
	Title   string
	StartAt time.Time
}

func getNextID() string {
	return uuid.New().String()
}

func NewEvent(title string, date string) (Event, error) {
	err := IsValidEvent(title, date)
	if err != nil {
		return Event{}, errors.New("Ошибка создания события:" + err.Error())
	}

	t, _ := dateparse.ParseAny(date)

	return Event{
		ID:      getNextID(),
		Title:   title,
		StartAt: t,
	}, nil
}
