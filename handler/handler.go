package handler

import (
	"main/service"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"main/model"
	"github.com/gin-gonic/gin"
)


func Create(c *gin.Context){
	input := model.UserInput{}

	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest , "Bad request")
		return
	}

	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()

	payload , err := user.Create(&input)
	if err != nil{
		log,Fatal(err)
	}

    c.JSON(http.StatusOK , payload)
}

func GetAll(c *gin.Context){
	service:= c.MustGet(service.ServiceKey).(service.Servicer)
	user:= service.NewUser()
	payload , err := user.GetAll()
	if err != nil{
	log.Fatal(err)
	}
	c.JSON(http.StatusOK , payload)
}

func GetOne (c *gin.Context){
    userID, err := strconv.Atoi((c.Param("user-id")))
	if err != nil{
		log.Fatal(err)
	}
	user = service.NewUser()
	payload, err := user.GetOne(userID)
	if err != nil{
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func Update(c *gin.Context){
	userID, err := strconv.Atoi((c.Param("user-id")))
	if err != nil{
		log.Fatal(err)
	}
	input := model.UserInput{}
	err = c.ShouldBindWith(&input, binding.JSON)
	if err != nil{
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()

	payload , err := user.Update(&input, userID)
	if err != nil{
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, payload)
}

func Delete (c *gin Context){
	userID, err := stconv.Atoi((c.Param("user-id"))) 
	if err != nil{
		log.Fatal(err)
	}
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()
	err = user.Delete(userID)
	if err != nil{
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, "削除されました")
}