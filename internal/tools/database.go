package tools

type DatabaseInterface interface {
	Create(bt BestTime) error
	GetAll() ([]BestTime, error)
}

type UserDatabaseInterface interface {
	CreateUser(user User) error
	GetUser(username string) (User, error)
}
