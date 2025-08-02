package events

import (
	"errors"
	"regexp"

	"github.com/araddon/dateparse"
)

func isValidTitle(title string) bool {
	pattern := "^[a-zA-Zа-яА-Я0-9 ,\\.]{3,50}$"

	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false
	}

	return matched
}

func IsValidEvent(title string, date string) error {
	if !isValidTitle(title) {
		return errors.New("неверный формат заголовка")
	}

	_, err := dateparse.ParseAny(date)
	if err != nil {
		return errors.New("неверный формат даты")
	}

	return nil
}
