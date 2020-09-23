package dao

import (
	"time"

	"doworksvc/pojo"
	"doworksvc/util/v"
	"doworksvc/util/xtime"
	"xorm.io/xorm"
)

type WorkDao struct {
	session *xorm.Session
}

func (dao *WorkDao) Add(item *pojo.Work) (*pojo.Work, error) {
	now := time.Now()
	item.TimeStamp = now.Unix()
	item.TimeView = now.Format(xtime.TIME_LAYOUT)
	_, err := dao.session.Insert(item)
	return item, err
}

func (dao *WorkDao) List(m v.M, page, size int) (res []pojo.Work, total int64, err error) {
	if m.Exist("completed") {
		dao.session = dao.session.Where("completed = ?", m.FBool("completed"))
	}
	if val := m.FInt64("start_time"); val > 0 {
		dao.session = dao.session.Where("belong_stamp>=?", val)
	}
	if val := m.FInt64("end_time"); val > 0 {
		dao.session = dao.session.Where("belong_stamp<=?", val)
	}
	total, err = dao.session.Desc("id").Limit(size, (page-1)*size).FindAndCount(&res)
	return
}

func (dao *WorkDao) Del(id uint64) error {
	_, err := dao.session.ID(id).Delete(new(pojo.Work))
	return err
}

func (dao *WorkDao) SetText(work *pojo.Work) error {
	_, err := dao.session.Table(new(pojo.Work)).ID(work.Id).Update(v.M{"text": work.Text})
	return err
}

func (dao *WorkDao) SetCompleted(id uint64, completed bool) error {
	_, err := dao.session.Table(new(pojo.Work)).ID(id).Update(v.M{"completed": completed})
	return err
}

func (dao *WorkDao) SetBelongDate(id uint64, date string) error {
	tm, err := xtime.ParseInLocal(xtime.TIME_LAYOUT_SIMPLE, date)
	if err != nil {
		return err
	}
	_, err = dao.session.Table(new(pojo.Work)).ID(id).Update(v.M{"belong_date": date, "belong_stamp": tm.Unix()})
	return err
}
