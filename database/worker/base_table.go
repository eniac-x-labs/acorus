package worker

import (
	"gorm.io/gorm"
)

type CreateTableDB interface {
	CreateTable(tableName, realTableName string)
}

type createTableDB struct {
	gorm *gorm.DB
}

func NewCreateTableDB(db *gorm.DB) CreateTableDB {
	return &createTableDB{gorm: db}
}
func (dao *createTableDB) CreateTable(tableName, realTableName string) {
	dao.gorm.Exec("CREATE TABLE IF NOT EXISTS " + tableName + " like " + realTableName)
}
