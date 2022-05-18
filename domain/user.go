package domain

type UserID string

func (self UserID) String() string {
	return string(self)
}

type User struct {
	ID             UserID
	Name, Password string
}

func NewUser(id UserID, name, password string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Password: password,
	}
}
