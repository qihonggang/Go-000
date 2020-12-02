package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func main() {
	service := UserService{}
	user, err := service.find("1")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows){
			log.Println(err)
			return
		}
	}
	fmt.Println(user)
}