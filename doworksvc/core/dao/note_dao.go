package dao

import (
	"doworksvc/pojo"
	"doworksvc/util/xtime"
	"time"
	"xorm.io/xorm"
)

type NoteDao struct {
	session *xorm.Session
}

func (dao *NoteDao) Add(item *pojo.Note) (*pojo.Note, error) {
	now := time.Now()
	item.TimeStamp = now.Unix()
	item.TimeView = now.Format(xtime.TIME_LAYOUT)
	_, err := dao.session.Insert(item)
	return item, err
}

func (dao *NoteDao) List(page, size int) (res []pojo.Note, total int64, err error) {
	total, err = dao.session.Desc("id").Limit(size, (page-1)*size).FindAndCount(&res)
	return
}

func (dao *NoteDao) Del(id uint64) error {
	_, err := dao.session.ID(id).Delete(new(pojo.Note))
	return err
}

func (dao *NoteDao) Update(item *pojo.Note) error {
	_, err := dao.session.ID(item.Id).Update(item)
	return err
}
