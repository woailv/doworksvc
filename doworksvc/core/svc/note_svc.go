package svc

import (
	"doworksvc/core/dao"
	"doworksvc/pojo"
)

type NoteSvc struct {
	*dao.DaoFactory
}

func (svc *NoteSvc) Add(item *pojo.Note) (*pojo.Note, error) {
	return svc.NoteDao().Add(item)
}

func (svc *NoteSvc) List(page, size int) ([]pojo.Note, int64, error) {
	return svc.NoteDao().List(page, size)
}

func (svc *NoteSvc) Del(id uint64) error {
	return svc.NoteDao().Del(id)
}

func (svc *NoteSvc) Update(item *pojo.Note) error {
	return svc.NoteDao().Update(item)
}
