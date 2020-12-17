package main

import (
	"database/sql"
	"errors"

	errs "github.com/pkg/errors"
)

type UserDao struct{}

type User struct {
	name string
}

func (user *UserDao) find(id string) (*User, error) {
	err := sql.ErrNoRows
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.Wrapf(err, "not found user : %s", id)
	}
	return &User{"qhg"}, nil
}
