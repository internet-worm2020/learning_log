package model

type User struct {
	Id     int
	Name   string
	Age    int
	Career string
}

func NewUser(id int, name string, age int, career string) *User {
	return &User{
		Id:     id,
		Name:   name,
		Age:    age,
		Career: career,
	}
}

func (u *User) ModifyUser(modifyUser *User) {
	if modifyUser.Name != "" {
		u.Name = modifyUser.Name
	}
	if modifyUser.Age != 0 {
		u.Age = modifyUser.Age
	}
	if modifyUser.Career != "" {
		u.Career = modifyUser.Career
	}
}