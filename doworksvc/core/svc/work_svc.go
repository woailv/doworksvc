package svc

import (
	"doworksvc/core/dao"
	"doworksvc/pojo"
	"doworksvc/util/v"
	"doworksvc/util/xtime"
)

type WorkSvc struct {
	*dao.DaoFactory
}

func (svc *WorkSvc) Add(item *pojo.Work) (*pojo.Work, error) {
	z := xtime.TodayZero()
	item.BelongDate = z.Format(xtime.TIME_LAYOUT_SIMPLE)
	item.BelongStamp = z.Unix()
	return svc.WorkDao().Add(item)
}

func (svc *WorkSvc) List(m v.M, page, size int) ([]pojo.Work, int64, error) {
	return svc.WorkDao().List(m, page, size)
}

func (svc *WorkSvc) Del(id uint64) error {
	return svc.WorkDao().Del(id)
}

func (svc *WorkSvc) SetText(work *pojo.Work) error {
	return svc.WorkDao().SetText(work)
}

func (svc *WorkSvc) SetCompleted(id uint64, completed bool) error {
	return svc.WorkDao().SetCompleted(id, completed)
}

func (svc *WorkSvc) SetBelongDate(id uint64, date string) error {
	return svc.WorkDao().SetBelongDate(id, date)
}
