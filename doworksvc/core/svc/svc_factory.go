package svc

import (
	"doworksvc/core/dao"
	"xorm.io/xorm"
)

func NewSvcFactory(session *xorm.Session) *SvcFactory {
	return &SvcFactory{daoFactory: dao.NewDaoFactory(session)}
}

type SvcFactory struct {
	daoFactory *dao.DaoFactory
}

func (svcFactory *SvcFactory) WorkSvc() *WorkSvc {
	return &WorkSvc{DaoFactory: svcFactory.daoFactory}
}

func (svcFactory *SvcFactory) UserSvc() *UserSvc {
	return &UserSvc{DaoFactory: svcFactory.daoFactory}
}

func (svcFactory *SvcFactory) NoteSvc() *NoteSvc {
	return &NoteSvc{DaoFactory: svcFactory.daoFactory}
}
