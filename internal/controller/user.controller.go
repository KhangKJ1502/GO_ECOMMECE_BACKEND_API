package controller

import (
	"go_ecommerce_backend/internal/service"
	"go_ecommerce_backend/pkg/respone"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc user controller
// us user service
// Controller => service => repo => models =>dbs
func (uc *UserController) GetUserByID(c *gin.Context) {
	// respone.SuccesRespone(c, 20003, []string{"khangKJ", "cr7", "KhangYT"}) //Nhận Repo không lỗilỗi
	respone.ErroRespone(c, 20003, "") // Nhận repo khi thiếu thông tin
}
