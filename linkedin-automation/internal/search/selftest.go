package search

import "errors"

func ValidateExtraction(profiles []string) error {
	if len(profiles) == 0 {
		return errors.New("no profiles extracted")
	}
	return nil
}
