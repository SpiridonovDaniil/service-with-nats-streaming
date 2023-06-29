package models

import "fmt"

var (
	ErrNotFound       = fmt.Errorf("user not found")
	ErrAlreadyInTheDB = fmt.Errorf("the user already exists in the database")
)
