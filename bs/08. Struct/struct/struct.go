package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {

	u := User{
		ID:   1,
		Name: "Alex",
	}
	fmt.Println(u)

	rename(u)
	fmt.Println(u)

	renameWithP(&u)
	fmt.Println(u)

	var zeroU User
	fmt.Println(zeroU)

	fmt.Println(u.FullName())

	u.SetName("САЛЕМ")

	u.SetNameWithP("КИРЮХА") // компилятор сам возьмёт &u
	fmt.Println(u)
}

// won't change source struct
func rename(u User) {
	u.Name = "Bob"
}

// will change source struct (because of pointer)
func renameWithP(u *User) {
	u.Name = "Bob"
}

// METHOD RECEIVER
func (u User) FullName() string {
	return u.Name
}

// VALUE RECEIVER (creates local variable of struct, won't change source)
func (u User) SetName(name string) {
	u.Name = name
}

// POINTER RECEIVER (change value with pointer)
func (u *User) SetNameWithP(name string) {
	u.Name = name
}

func compisitionExample() {

	type Profile struct {
		Age int
	}

	type User struct {
		Name    string
		Profile Profile
	}

}

func aggregatedStructures() {
	type Logger struct {
		Level string
	}

	type Service struct {
		Logger
	}

	var s Service

	fmt.Println(s.Level)
	fmt.Println(s.Logger.Level)
}
