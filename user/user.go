package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: "Arshad",
			LastName:  "Jamal",
			BirthDate: "20/01/1999",
			CreatedAt: time.Now(),
		},
	}
}
func (u *User) OutputUserDet() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func (u *User) Clearoutput() {
	u.FirstName = ""
	u.FirstName = ""
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name and birth date are required")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	}, nil
}
