package service

import (
	"main/model"
	"time"
	"xorm.io/xorm"
)


type Users struct{
	engine *xorm.Engine
}

func NewUsers(engine *xorm.Engine) *Users {
	u := Users{
		engine: engine,
	}
	return &u
}

func (u* Users) Create(input *model.UserInput)(*model.Users , error){
	now := time.Now()
	user := model.Users{
		Name: input.Name,
		Address: input.Address,
		CreatedAt: now,
		UpdatedAt: now,
	}
	_, err:= u.engine.Table("users").Insert(&user)
	if err != nil{
		return nil, err
	}
	return &user , nil
}

