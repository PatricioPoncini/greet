package greet

import (
	"fmt"
)

// User represents a simple user with a name.
type User struct {
	Name string
}

// SayHi returns a greeting message that includes the user's name.
func (u User) SayHi() string {
	return fmt.Sprintf("Hi, %s!", u.Name)
}
