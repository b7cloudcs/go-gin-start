package service

import (
	"errors"
)

type User struct{}

/**
 * Service Demo
**/
func (User) TheDemo() (string, error) {
	return "Service Demo", errors.New("")
}