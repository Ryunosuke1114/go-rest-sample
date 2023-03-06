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

func (u *Users) GetOne(userID int)(*model.Users, error){
	user := model.Users{}
	_, err := u.engine.Table("users").Where("id = ?", userID).Get(&user)
	if err != nil{
		return nil, err
	}
	return &user, nil
}

func (u *Users) GetAll()([]*model.Users , error){
	users := []*model.Users{}
	err := u.engine.Table("users").Find(&users)
	if err != nil{
		return nil, err
	}
	return users , nil
}

func (u *Users) Update(input *model.UserInput , userID int)(*model.Users, error){
	user := model.Users{
		Name: input.Name,
		Address: input.Address,
	}
	_ , err := u.engine.Table("users").Where("id = ?" , userID).Update(&user)
	if err != nil{
		return nil, err
	}
	return u.GetOne(userID)
}

func (u *Users) Delete (userID int) error {
	user := model.Users{}
	_, err := u.engine.Table("users").Where("id = ?", userID).Delete(&user)
	if err != nil{
		return err
	}
	return nil
}

