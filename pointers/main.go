package main

import "fmt"

// when are we going to use pointers
// 1. when we need to update state
// 2. when we want to optimize memory for large objects that are getting called A LOT

type User struct {
	email string
	username string
	age int
}

func (user User) Email() string {
	return user.email
}

func (user *User) updateEmail(email string) {
	user.email = email
}

func (user User) Username() string {
	return user.username
}

func (user *User) updateUsername(username string) {
	user.username = username
}

func (user User) Age() int {
	return user.age
}

func (user *User) updateAge(age int) {
	user.age = age
}

func main() {
	user := User {
		email: "agg@foo.com",
		username: "agg",
		age: 21,
	}

	user.updateEmail("foo@bar.com")
	user.updateUsername("foo")
	user.updateAge(22)

	fmt.Println(user.Email())
	fmt.Println(user.Username())
	fmt.Println(user.Age())
}