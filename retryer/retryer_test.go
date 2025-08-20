package retryer_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/felipe-bonato/goutils/retryer"
)

var ErrGeneric error = errors.New("failed")

func TestRetryerFailed(t *testing.T) {
	err := retryer.Retry(5, func(try int) error {
		return fmt.Errorf("%w: on try %d", ErrGeneric, try)
	})
	if !errors.Is(err, ErrGeneric) {
		t.Fatalf("Expected %v, but got %v", ErrGeneric, err)
	}

	if err.Error() != "failed: on try 5" {
		t.Fatalf("Expected to err on try 5, but got %v", err)
	}
}

func TestRetryerSuccess(t *testing.T) {
	err := retryer.Retry(5, func(try int) error {
		return nil
	})
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
}

func TestRetryerSuccessAfterRetries(t *testing.T) {
	err := retryer.Retry(5, func(try int) error {
		if try == 5 {
			return nil
		}

		return ErrGeneric
	})
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
}
