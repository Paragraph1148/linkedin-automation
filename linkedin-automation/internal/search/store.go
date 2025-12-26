package search

import (
	"encoding/json"
	"os"
)

func SaveProfiles(profiles []string, path string) error {
	data, err := json.MarshalIndent(profiles, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}