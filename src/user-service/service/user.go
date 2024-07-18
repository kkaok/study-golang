package service

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserService interface {
	CreateUser(User) (string, error)
	GetUser(string) (User, error)
	GetUsers() ([]User, error)
	UpdateUser(User) error
	DeleteUser(string) error
}
