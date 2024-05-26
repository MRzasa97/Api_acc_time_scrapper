package tools

type BestTime struct {
	CarModel  string `json:"car_model"`
	BestTime  string `json:"best_time"`
	TrackName string `json:"track_name"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"json:password"`
}
