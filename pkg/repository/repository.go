package repository

type Repository interface {
	Users
}

type Users interface {
	CreateUser(user User) error
	ReadUser(login string) ([]User, error)
	ReadUserByEmail(email string) ([]User, error)
	UpdateUser(newUser User) error
	DeleteUser(user User) error
}
