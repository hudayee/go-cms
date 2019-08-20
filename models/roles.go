package models

import (
	"time"
)

type Role struct {
	Id        int        `xorm:"not null pk autoincr INTEGER" json:"id" form:"id"`
	Name      string     `xorm:"not null unique VARCHAR(10)" json:"name"  form:"name"`
	Status    string     `xorm:"not null VARCHAR(10)" json:"status"  form:"status"`
	CraetedAt time.Time  `xorm:"not null TIMESTAMPZ created" json:"createdAt"  form:"createdAt"`
	UpdatedAt time.Time  `xorm:"not null TIMESTAMPZ updated" json:"updatedAt"  form:"updatedAt"`
	DeletedAt *time.Time `xorm:"TIMESTAMPZ deleted" json:"deletedAt"  form:"deletedAt"`
}

type roleDao struct {
}
type RoleQuery struct {
	Id     int
	Name   string
	Status string
	Offset int
	Limit  int
}

var RoleDao *roleDao

func (dao *roleDao) Get(id int) (role Role, err error) {
	role.Id = id
	_, err = Engine.Table("roles").Get(&role)
	return
}

func (dao *roleDao) GetAll(query RoleQuery) (roles []*Role, count int64, err error) {
	roles = make([]*Role, 0)
	engine := Engine.Table("roles").Alias("r")
	if query.Id != 0 {
		engine.Where("r.id=?", query.Id)
	}
	if query.Name != "" {
		engine.Where("r.name ilike ?", "%"+query.Name+"%")
	}
	count, err = engine.Limit(query.Limit, query.Offset).FindAndCount(&roles)
	return
}

func (dao *roleDao) Create(role Role) error {
	_, err := Engine.Table("roles").InsertOne(role)
	return err
}

func (dao *roleDao) Update(role Role) (int64, error) {
	//mustCols []string   .Cols(mustCols...)
	return Engine.Table("roles").ID(role.Id).Update(role)
	//return errors.New(fmt.Sprintf("not found roleId=%d ", role.Id))

}
