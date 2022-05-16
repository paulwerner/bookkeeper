package domain

type UserID string

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
