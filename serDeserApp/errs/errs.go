package errs

import "fmt"

func WrapIfErr(msg string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", msg, err)
	}
	return nil
}
