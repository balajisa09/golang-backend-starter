package models


type Album struct{
	ID string `json:"id" gorm:"primary key"`
	Title string `json:"title"`
	Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
