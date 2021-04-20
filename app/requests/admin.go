package requests

type AdminRegisterRequest struct {
	Username string `form:"username" json:"username" validate:"required" label:"用户名"`
	Password string `form:"password" json:"password" validate:"required" label:"密码"`
}
