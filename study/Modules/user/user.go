package user

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) User {
	if name == "" {
		return User{}
	}
	if age <= 0 && age >= 150 {
		return User{}
	}

	return User{
		name: name,
		age:  age,
	}
}

func (u User) SetName(name string) User {
	if name == "" {
		return u
	}

	u.name = name
	return u
}

func (u User) SetAge(age int) User {
	if age <= 0 && age >= 150 {
		return u
	}

	u.age = age
	return u
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetAge() int {
	return u.age
}
