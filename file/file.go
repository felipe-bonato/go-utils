package file

import (
	"errors"
	"fmt"
	"os"
)

// Returns true if exists, false if it doesn't, and false and error if unknowable.
func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		// It's only guaranteed that the file doesn't exists if the return is ErrNotExist.
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		// Schrodinger file: it may or may not exist.
		// See for more info: https://stackoverflow.com/a/12518877
		return false, fmt.Errorf("cannot know if file exists: %w", err)
	}

	return true, nil
}
