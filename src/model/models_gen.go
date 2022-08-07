// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

// can only be used if the user is logged in
type NewTodo struct {
	Text string `json:"text"`
	// Are you done?
	Done *bool `json:"done"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleOwner Role = "OWNER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleOwner,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleOwner:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}