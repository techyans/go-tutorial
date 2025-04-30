package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *User) OutputUserDet() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) Clearoutput() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthDate string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

func main() {
	firstName := getUserData("Please enter your first name")
	lastName := getUserData("Please enter your last name")
	birthDate := getUserData("Please enter your birth date (MM/DD/YYYY)")

	var Appuser User

	// Appuser = User{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	birthDate: birthDate,
	// 	createdAt: time.Now(),
	// }
	Appuser = *newUser(firstName, lastName, birthDate)

	Appuser.OutputUserDet()
	Appuser.Clearoutput()
	Appuser.OutputUserDet()

}

func getUserData(userdata string) string {
	fmt.Println(userdata)
	var value string
	fmt.Scan(&value)
	return value
}
