package user

import (
	"fmt"
)

type User struct {
	ID   string
	Name string
}

func SignUp(Name string) (*User, error) {
	if Name == "" {
		return nil, fmt.Errorf("invalid parameters")
	}

	return &User{
		ID:   "generated_id",
		Name: Name,
	}, nil
}
