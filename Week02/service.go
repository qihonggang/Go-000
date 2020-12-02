package main

type UserService struct {}

func (user *UserService) find(id string) (*User, error) {
	dao := UserDao{}
	return dao.find(id)
}