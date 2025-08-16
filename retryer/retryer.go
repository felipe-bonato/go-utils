package retryer

// Runs ´f´ up to ´retries´ times if it returns error.
func Retry(retries int, f func(try int) error) error {
	var err error

	for try := 1; try <= retries; try++ {
		if err = f(try); err == nil { // If NO error
			return nil
		}
	}

	return err
}

// Same as ´Retry´, but call ´onErr´ every time an error happens,
// and if it returns error, the entire ´RetryErr´ will return that error.
func RetryErr(retries int, f func(try int) error, onErr func(try int, err error) error) error {
	var err error

	for try := 1; try <= retries; try++ {
		if err = f(try); err == nil { // If NO error
			return nil
		}
		if err = onErr(try, err); err != nil {
			return err
		}
	}

	return err

}
