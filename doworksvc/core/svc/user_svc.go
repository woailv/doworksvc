package svc

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"doworksvc/conf"
	"doworksvc/core/dao"
	"doworksvc/pojo"
	"doworksvc/util/exmap"
)

func NewUser(factory *dao.DaoFactory) *UserSvc {
	return &UserSvc{DaoFactory: factory}
}

type UserSvc struct {
	*dao.DaoFactory
}

func (s *UserSvc) Register(param *pojo.UserRegisterParam) (*pojo.User, error) {
	//code, ok := exmap.Get(conf.GenUserRegisterEmailKey(param.Email))
	//if !ok {
	//	return nil, fmt.Errorf("验证码:%s,不存在", param.Code)
	//} else if code.(string) != param.Code {
	//	return nil, errors.New("验证码错误")
	//}
	ok, err := s.UserDao().ExistEmail(param.Email)
	if err != nil {
		return nil, err
	} else if ok {
		return nil, fmt.Errorf("邮箱:%s,已被注册", param.Email)
	}
	return s.UserDao().Add(&pojo.User{
		Uid:   param.Uid,
		Pwd:   param.Pwd,
		Email: param.Email,
	})
}

func (s *UserSvc) Login(param *pojo.UserLoginParam) (*pojo.User, error) {
	user, exist, err := s.UserDao().GetByEmail(param.Email)
	if err != nil {
		return nil, err
	} else if !exist {
		return nil, errors.New("邮箱不存在")
	}
	if user.Pwd != param.Pwd {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

func (s *UserSvc) UserRegisterEmailCode(toEmail string) error {
	rand.Seed(time.Now().Unix())
	code := fmt.Sprintf("%03v", rand.Intn(999))
	log.Printf("邮箱:%s,验证码:%s", toEmail, code)
	//err := email.Send("dowork注册验证码", fmt.Sprintf("你的注册验证码为:%s,5分钟内有效", code), toEmail)
	//if err != nil {
	//	return err
	//}
	exmap.Set(conf.GenUserRegisterEmailKey(toEmail), code, 300)
	return nil
}
