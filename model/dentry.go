package model

import (
	"gorm.io/gorm"
	"time"
)

const DentryTableName = "dentrys"

type Dentry struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement"`
	ParentID    int       `gorm:"column:parent_id;index:parentId_idx;uniqueIndex:parent_dentry_idx"`
	Name        string    `gorm:"column:name;index:dentry_name_idx;uniqueIndex:parent_dentry_idx"`
	Description string    `gorm:"column:description"`
	NodeId      int       `gorm:"column:node_id"`
	Type        string    `gorm:"column:type"`
	CreatedAt   time.Time `gorm:"column:create_at"`
	UpdatedAt   time.Time `gorm:"column:update_at"`
}

func (g *Dentry) TableName() string {
	return DentryTableName
}

func GetDentryByParentID(db *gorm.DB, parentId int) (dentrys []Dentry) {
	db.Table(DentryTableName).Where("parent_id = ?", parentId).Find(&dentrys)
	return
}

func GetDentryIdByNameAndParentID(db *gorm.DB,name string, parentId int) (dentry Dentry,err error) {
	err = db.Table(DentryTableName).Where("name = ? and parent_id = ?",name, parentId).First(&dentry).Error
	return
}