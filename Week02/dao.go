package main

import (
	"database/sql"
	errs "github.com/pkg/errors"
	"errors"
)

type UserDao struct {}

type User struct {
	name string
}

func (user *UserDao)find(id string) (*User, error) {
	err := sql.ErrNoRows
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.Wrapf(err, "Not found id : %s", id)
	}
	return &User{"qhg"}, nil
}