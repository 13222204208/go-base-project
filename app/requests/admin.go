package requests

type AdminRegisterRequest struct {
	Username string `form:"username" json:"username" validate:"required" label:"用户名"`
	Password string `form:"password" json:"password" validate:"required" label:"密码"`
	Name     string `form:"name" json:"name" validate:"required" label:"姓名"`
	Phone    string `form:"phone" json:"phone" validate:"required" label:"手机号"`
}

//LoginRequest 登陆入参
type AdminLoginRequest struct {
	Username string `form:"username" json:"username" validate:"required" label:"用户名"`
	Password string `form:"password" json:"password" validate:"required" label:"密码"`
}

type GoodsTypeRequest struct {
	Name string `form:"name" json:"name" validate:"required" label:"分类名称"`
}
