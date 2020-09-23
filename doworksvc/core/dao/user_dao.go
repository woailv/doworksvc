package dao

import (
	"doworksvc/pojo"
	"errors"
	"strings"
	"xorm.io/xorm"
)

type UserDao struct {
	session *xorm.Session
}

func (dao *UserDao) Add(item *pojo.User) (*pojo.User, error) {
	_, err := dao.session.Insert(item)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, errors.New("用户名已存在")
		}
	}
	return item, err
}

func (dao *UserDao) List(page, size int) (res []pojo.User, total int64, err error) {
	total, err = dao.session.Desc("id").Limit(size, (page-1)*size).FindAndCount(&res)
	return
}

func (dao *UserDao) Del(id uint64) error {
	_, err := dao.session.ID(id).Delete(new(pojo.User))
	return err
}

func (dao *UserDao) GetByUid(uid string) (*pojo.User, bool, error) {
	res := new(pojo.User)
	exist, err := dao.session.Where("uid = ?", uid).Get(res)
	return res, exist, err
}

func (dao *UserDao) GetByEmail(email string) (*pojo.User, bool, error) {
	res := new(pojo.User)
	exist, err := dao.session.Where("email = ?", email).Get(res)
	return res, exist, err
}

func (dao *UserDao) ExistEmail(email string) (bool, error) {
	count, err := dao.session.ForUpdate().Where("email=?", email).Count(new(pojo.User))
	return count > 0, err
}
