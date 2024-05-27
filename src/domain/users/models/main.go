package models

type User struct {
	id int
}

func (model User) GetId() int {
	return model.id
}

func NewUser(id int) User {
	return User{id: id}
}
