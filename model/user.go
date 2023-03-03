package model

import "time"
type Users struct {
	ID        int       `xorm:"id pk autoincr"`
	Name      string    `xorm:"name"`
	Address   string    `xorm:"address"`
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
}

type UserInput struct {
	Name string `json:"name" binding:"required" xorm:"name"`	
	Address string `json:"address" binding:"required" xorm:"address"`
}