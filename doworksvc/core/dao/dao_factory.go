package dao

import (
	"fmt"

	"xorm.io/xorm"
)

func NewDaoFactory(db *xorm.Session) *DaoFactory {
	return &DaoFactory{db}
}

type DaoFactory struct {
	session *xorm.Session
}

func (daoFactory *DaoFactory) TxExe(f func() error) error {
	err := daoFactory.session.Begin()
	if err != nil {
		return fmt.Errorf("开启事务失败:%s", err)
	}
	err = f()
	if err != nil {
		_ = daoFactory.session.Rollback()
		return err
	}
	err = daoFactory.session.Commit()
	if err != nil {
		return fmt.Errorf("提交事务失败:%s", err)
	}
	return nil
}

func (daoFactory *DaoFactory) WorkDao() *WorkDao {
	return &WorkDao{daoFactory.session}
}

func (daoFactory *DaoFactory) UserDao() *UserDao {
	return &UserDao{daoFactory.session}
}

func (daoFactory *DaoFactory) NoteDao() *NoteDao {
	return &NoteDao{daoFactory.session}
}
