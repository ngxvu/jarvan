package model

type Item struct {
	Url string `json:"url" gorm:"primaryKey"`
}
