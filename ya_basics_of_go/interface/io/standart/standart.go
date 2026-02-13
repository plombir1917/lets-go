package standart

import "fmt"

type User struct {
	Email string
	Pwd   string
}

func (u User) String() string {
	return "user with email " + u.Email
}

func Execute() {
	u := User{Email: "example@mail.com"}

	fmt.Println(u)
}
