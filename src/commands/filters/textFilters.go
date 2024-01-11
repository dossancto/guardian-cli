package filters

import (
	"errors"
	"regexp"
)

func SpecialCharFilter(text string) error {
	reg := `[^a-zA-Z0-9]`
	onlyChars := regexp.MustCompile(reg).MatchString(text)

	if onlyChars {
		return errors.New("The text must match '" + reg + "'")
	}

	return nil
}

func BlankTextFilter(text string) error {
	if text == "" {
		return errors.New("Blank Text")
	}

	return nil
}
