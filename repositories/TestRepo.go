package repositories

import (
	"github.com/jinzhu/gorm"
	"json2struct/models"
)

type TestRepositories struct {
	db *gorm.DB
}

func NewTestRepositories() *TestRepositories {
	return &TestRepositories{db: models.DB.Mysql}
}

func (t *TestRepositories) Create(userAgent, Ip, Addr, rotateStr string) bool {
	var rotateRecord models.RotateCountRecord
	rotateRecord.Address = Addr
	rotateRecord.Ip = Ip
	rotateRecord.UserAgent = userAgent
	rotateRecord.RotateStr = rotateStr
	if err := t.db.Create(&rotateRecord).Error; err != nil {
		return false
	}
	return true
}

func (t *TestRepositories) GetRotateCount(count *int) {
	var records []*models.RotateCountRecord
	t.db.Find(&records)
	*count = len(records)
}
