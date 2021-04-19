package web

//验证接口
type Validator interface {
	Validate(interface{}) (bool, error)
}

type MyTest struct {
}

func (v MyTest) AdminLogin(val interface{}) (bool, error) {
	return true, nil
}
