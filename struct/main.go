package main

import (
	"fmt"
	"go_course/user"
)

func main() {
	firstName := getUserData("Please enter your first name")
	lastName := getUserData("Please enter your last name")
	birthDate := getUserData("Please enter your birth date (MM/DD/YYYY)")

	var Appuser *user.User

	// Appuser = User{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	birthDate: birthDate,
	// 	createdAt: time.Now(),
	// }
	admin := user.NewAdmin("admin", "admin@123")
	admin.OutputUserDet()
	admin.Clearoutput()
	admin.OutputUserDet()

	Appuser, err := user.NewUser(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
	}
	Appuser.OutputUserDet()
	Appuser.Clearoutput()
	Appuser.OutputUserDet()

}

func getUserData(userdata string) string {
	fmt.Println(userdata)
	var value string
	fmt.Scanln(&value)
	return value
}
