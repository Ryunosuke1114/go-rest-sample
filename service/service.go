package service

import (
	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

type Service struct {
	engine *xorm.Engine
}

func NewService(engine *xorm.Engine) Servicer{
	return &Service{engine}
}
type Servicer interface {
	NewUser() *Users
}
func (s *Service) NewUser() *Users{
	return NewUsers(s.engine)
}

const (
	ServiceKey = "service_factory"
)

func  ServiceFactoryMiddleware(svc Servicer) gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Set(ServiceKey , svc)
		c.Next()
	}
}