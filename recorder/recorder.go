package recorder

import (
	"errors"
	"fmt"

	"github.com/rockiecn/backend2/events"
)

// parser for register account
func ReAccRecorder(e interface{}) error {
	// assertion
	v, ok := e.(*events.ReAcc)
	if !ok {
		return errors.New("assertion failed")
	}

	fmt.Println("output event struct:")
	fmt.Println(v)

	// todo: record reAcc into db

	return nil
}
