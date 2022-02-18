package model

type User struct {
	Id       int    `json:"id" binding:"required" gorm:"unique;notNull"`
	UserName string `json:"name" binding:"required" gorm:"primaryKey"`
	Score    int    `json:"score" gorm:"default:null"`
}
