package dal

type userDal struct {}

var UserDal = userDal{}

func (userDal) GetUserId(userName, password string) int64 {
	if userName == "hzh" && password == "hzh" {
		return 1
	}
	return 0
}