package dal

type userDal struct {}

var UserDal = userDal{}

func (userDal) GetUserId(userName, password string) int64 {
	if userName == "hzh" && password == "hzh" {
		return 1
	}
	return 0
}

func (userDal) NewUser(userName, password string) int64 {
	if UserDal.GetUserId(userName, password) != 0 {     // 账号已存在
		return 0
	}
	// todo 保存至数据库
	return 2
}
