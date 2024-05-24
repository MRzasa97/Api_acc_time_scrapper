package tools

type BestTime struct {
	CarModel string `json:"car_model"`
	BestTime string `json:"best_time"`
}

type DatabaseInterface interface {
	Create(bt BestTime) error
	GetAll() ([]BestTime, error)
}
