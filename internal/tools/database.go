package tools

type DatabaseInterface interface {
	Create(BestTime, int) error
	GetAll() ([]BestTime, error)
	GetTrack(string) (Track, error)
	CreateTrack(string) (Track, error)
}

type UserDatabaseInterface interface {
	CreateUser(user User) error
	GetUser(username string) (User, error)
}
